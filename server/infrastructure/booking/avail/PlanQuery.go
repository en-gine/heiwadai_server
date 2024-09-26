package avail

import (
	"errors"
	"strconv"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/util"
	"server/infrastructure/env"
	"server/infrastructure/logger"
)

var _ queryservice.IPlanQueryService = &PlanQuery{}
var (
	TLBookingUser      = env.GetEnv(env.TlbookingUsername)
	TLBookingPass      = env.GetEnv(env.TlbookingPassword)
	TLBookingSearchURL = env.GetEnv(env.TlbookingAvailApiUrl)
)

type PlanQuery struct {
	storeQuery queryservice.IStoreQueryService
}

func NewPlanQuery(storeQuery queryservice.IStoreQueryService) *PlanQuery {
	return &PlanQuery{
		storeQuery: storeQuery,
	}
}

func (p *PlanQuery) Search(
	stores []*entity.StayableStore,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes []entity.SmokeType,
	mealType entity.MealType,
	roomTypes []entity.RoomType,
) (*[]entity.PlanCandidate, error) {
	type Result struct {
		candidates *[]entity.PlanCandidate
		err        error
	}
	nights := stayTo.Sub(stayFrom).Hours() / 24
	if nights > 99 {
		return nil, errors.New("99泊以上の予約はできません")
	}
	personTotal := adult + child
	if personTotal > 99 {
		return nil, errors.New("99名以上の予約はできません")
	}
	chLength := len(stores) * len(roomTypes)
	resultsCh := make(chan Result, chLength)
	for _, store := range stores {
		var bookingID string
		if env.GetEnv(env.TlbookingIsTest) == "true" {
			bookingID = "E69502"
		} else {
			bookingID = store.BookingSystemID
		}
		for _, roomType := range roomTypes {
			go func(st *entity.StayableStore, rt entity.RoomType) {
				reqBody := NewOTAHotelAvailRQ(
					bookingID,
					stayFrom,
					stayTo,
					adult,
					child,
					roomCount,
					smokeTypes,
					mealType,
					rt,
				)
				request := NewEnvelopeRQ(*st, reqBody)
				res, err := util.Request[EnvelopeRQ, EnvelopeRS](TLBookingSearchURL, request)
				if err != nil {
					resultsCh <- Result{nil, err}
					return
				}

				if res.Body.OTA_HotelAvailRS.Errors != nil {
					errs := res.Body.OTA_HotelAvailRS.Errors
					msg := errs.Error[0].ShortText
					logger.Error(msg)
					resultsCh <- Result{nil, errors.New(msg)}
					return
				}
				guestCount := adult + child

				candidates, err := p.AvailRSToCandidates(res, roomCount, guestCount, int(nights))
				if err != nil {
					logger.Error(err.Error())
					resultsCh <- Result{nil, err}
					return
				}
				resultsCh <- Result{candidates, nil}
			}(store, roomType)
		}
	}

	var allCandidates []entity.PlanCandidate
	for i := 0; i < chLength; i++ {
		res := <-resultsCh
		if res.err != nil {
			return nil, res.err
		}
		allCandidates = append(allCandidates, *res.candidates...)
	}
	return &allCandidates, nil
}

func (p *PlanQuery) GetPlanDetailByID(
	planID string,
	store *entity.StayableStore,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	TlBookingRoomTypeCode string,
) (*entity.PlanStayDetail, error) {
	var hotelCode string
	if env.GetEnv(env.TlbookingIsTest) == "true" {
		hotelCode = "E69502"
	} else {
		hotelCode = store.StayableStoreInfo.BookingSystemID
	}

	reqBody := NewOTAHotelPlanDetailRQ(
		planID,
		hotelCode,
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		TlBookingRoomTypeCode,
	)
	request := NewEnvelopeRQ(*store, reqBody)
	res, err := util.Request[EnvelopeRQ, EnvelopeRS](TLBookingSearchURL, request)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	guestCount := adult + child
	plan, err := p.AvailDetailRSToPlanDetail(res, roomCount, guestCount)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return plan, nil
}

func (p *PlanQuery) AvailRSToCandidates(res *EnvelopeRS, roomCount int, guestCount int, nights int) (*[]entity.PlanCandidate, error) {
	var candidates []entity.PlanCandidate
	body := res.Body.OTA_HotelAvailRS

	for _, roomStay := range body.RoomStays.RoomStay {
		hotelCode := roomStay.RPH
		var stayable *entity.StayableStore
		var err error
		if env.GetEnv(env.TlbookingIsTest) == "true" {
			stayables, err := p.storeQuery.GetStayables()
			if err != nil {
				return nil, err
			}
			stayable = stayables[0]
		} else {
			stayable, err = p.storeQuery.GetStayableByBookingID(hotelCode)
		}
		if err != nil {
			return nil, err
		}
		if len(roomStay.RoomTypes.RoomType) == 0 { // room not found
			return &[]entity.PlanCandidate{}, nil
		}

		RoomTypeObject := roomStay.RoomTypes.RoomType[0]
		apiRoomTypeCode := RoomTypeObject.RoomTypeCode
		apiRoomTypeName := RoomTypeObject.RoomDescription.Name
		entityRoomType := BedTypeCodeToRoomType(RoomTypeObject.BedTypeCode)
		smokeType := IsNonSmokingToSmokeType(RoomTypeObject.NonSmoking) // true false nil

		for index, plan := range roomStay.RatePlans.RatePlan {
			// 一泊毎や人数ごとの追加料金
			tmpAmount := roomStay.RoomRates.RoomRate[index].Total.AmountAfterTax
			var planPrice uint64
			roomPrice, _ := strconv.ParseUint(tmpAmount, 10, 64)

			planPrice = roomPrice * uint64(roomCount)
			availStatus := AvailabilityStatus(roomStay.RoomRates.RoomRate[index].AvailabilityStatus)
			if availStatus == AvailableClosedOut {
				//　売り切れ
				continue
			}
			planID := plan.RatePlanCode
			planName := plan.RatePlanName

			var planImageURL string = ""
			if len(plan.RatePlanDescription.URL.Value) > 0 {
				planImageURL = plan.RatePlanDescription.URL.Value
			}

			var planOverView string = ""
			if len(plan.RatePlanDescription.Text.Value) > 0 {
				planOverView = plan.RatePlanDescription.Text.Value
			}

			var IncludeBreakfast bool = *plan.MealsIncluded.Breakfast
			var IncludeDinner bool = *plan.MealsIncluded.Dinner

			plan := entity.RegenPlan(
				planID,
				planName,
				uint(planPrice),
				planImageURL,
				entityRoomType,
				entity.MealType{
					Morning: IncludeBreakfast,
					Dinner:  IncludeDinner,
				},
				smokeType,
				planOverView,
				stayable.ID,
				apiRoomTypeCode,
				apiRoomTypeName,
			)

			candidate := entity.NewPlanCandidate(plan, nights, guestCount)

			candidates = append(candidates, *candidate)
		}
	}
	return &candidates, nil
}

func (p *PlanQuery) AvailDetailRSToPlanDetail(res *EnvelopeRS, roomCount int, guestCount int) (*entity.PlanStayDetail, error) {
	var plans []entity.Plan
	body := res.Body.OTA_HotelAvailRS
	var planStayDateInfos []entity.StayDateInfo
	hotelCode := body.Criteria.Criterion.HotelRef.HotelCode
	for _, roomStay := range body.RoomStays.RoomStay {
		var stayable *entity.StayableStore
		var err error
		if env.GetEnv(env.TlbookingIsTest) == "true" {
			stayables, err := p.storeQuery.GetStayables()
			if err != nil {
				return nil, err
			}
			stayable = stayables[0]
		} else {
			stayable, err = p.storeQuery.GetStayableByBookingID(hotelCode)
		}
		if err != nil {
			return nil, err
		}
		if len(roomStay.RoomTypes.RoomType) == 0 { // room not found
			return &entity.PlanStayDetail{}, nil
		}

		roomTypeObject := roomStay.RoomTypes.RoomType[0]
		entityRoomType := BedTypeCodeToRoomType(roomTypeObject.BedTypeCode)

		smokeType := IsNonSmokingToSmokeType(roomTypeObject.NonSmoking) // true false nil

		availStatus := AvailabilityStatus(roomStay.RoomRates.RoomRate[0].AvailabilityStatus)
		if availStatus == AvailableClosedOut {
			//　売り切れ
			continue
		}

		var planTotalPrice uint
		// 合計金額の計算
		for _, room := range roomStay.RoomRates.RoomRate {
			// 一泊毎や人数ごとの追加料金
			var planStayDateInfo entity.StayDateInfo
			stayDate, err := room.EffectiveDate.ToDate()
			if err != nil {
				return nil, errors.New("EffectiveDateの変換に失敗しました")
			}
			planStayDateInfo.StayDate = stayDate

			roomPrice, _ := strconv.ParseUint(room.Total.AmountAfterTax, 10, 64)
			dayTotalPrice := (uint(roomPrice) * uint(roomCount))
			planStayDateInfo = entity.StayDateInfo{
				StayDate:           stayDate,
				StayDateTotalPrice: dayTotalPrice,
			}

			planTotalPrice = planTotalPrice + dayTotalPrice
			planStayDateInfos = append(planStayDateInfos, planStayDateInfo)
		}

		for index, plan := range roomStay.RatePlans.RatePlan {

			availStatus := AvailabilityStatus(roomStay.RoomRates.RoomRate[index].AvailabilityStatus)
			if availStatus == AvailableClosedOut {
				//　売り切れ
				return &entity.PlanStayDetail{}, nil
			}
			planID := plan.RatePlanCode
			planName := plan.RatePlanName

			var planImageURL string = ""
			if len(plan.RatePlanDescription.URL.Value) > 0 {
				planImageURL = plan.RatePlanDescription.URL.Value
			}

			var planOverView string = ""
			if len(plan.RatePlanDescription.Text.Value) > 0 {
				planOverView = plan.RatePlanDescription.Text.Value
			}

			var IncludeBreakfast bool = *plan.MealsIncluded.Breakfast
			var IncludeDinner bool = *plan.MealsIncluded.Dinner

			plan := entity.RegenPlan(
				planID,
				planName,
				planTotalPrice,
				planImageURL,
				entityRoomType,
				entity.MealType{
					Morning: IncludeBreakfast,
					Dinner:  IncludeDinner,
				},
				smokeType,
				planOverView,
				stayable.ID,
				roomTypeObject.RoomTypeCode,
				roomTypeObject.RoomDescription.Name,
			)

			plans = append(plans, *plan)
		}
	}
	if len(plans) == 0 {
		return &entity.PlanStayDetail{}, nil
	}
	planDetail := &entity.PlanStayDetail{
		Plan:          &(plans)[0],
		StayDateInfos: &planStayDateInfos,
	}
	return planDetail, nil
}

func NewOTAHotelAvailRQ(
	hotelCode string,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes []entity.SmokeType,
	mealType entity.MealType,
	roomType entity.RoomType,
) *OTA_HotelAvailRQ {
	// 日付
	start := util.DateToYYYYMMDD(stayFrom)
	end := util.DateToYYYYMMDD(stayTo)

	// ホテルコード
	var hotelRef HotelRef = HotelRef{HotelCode: hotelCode}
	mealsIncluded := MealTypeToQuery(mealType)
	bedTypeCode := RoomTypeToBedType(roomType)
	roomStayCandidate := NewRoomStayCandidate(
		nil,         // RoomTypeCode
		bedTypeCode, // BedTypeCode
		adult,
		&child,
		roomCount,
		smokeTypes,
		nil, // EffectiveDate
		nil, // ExpireDate,
	)

	pricingMethod := PricingMethodLowestperstay
	return &OTA_HotelAvailRQ{
		Version:       "1.0",
		PrimaryLangID: "jpn",
		HotelStayOnly: util.BoolPtr(true), // ホテル情報のみを返すフラグ。trueにしないと返ってこない
		PricingMethod: &pricingMethod,
		AvailRequestSegments: AvailRequestSegments{
			AvailRequestSegment: AvailRequestSegment{
				HotelSearchCriteria: HotelSearchCriteria{
					Criterion: Criterion{
						HotelRef: hotelRef,
						RatePlanCandidates: &RatePlanCandidates{ // 食事タイプ
							RatePlanCandidate: []RatePlanCandidate{
								{
									MealsIncluded: mealsIncluded,
								},
							},
						},
						StayDateRange: &StayDateRange{
							Start: &start,
							End:   &end,
						},
						RoomStayCandidates: &RoomStayCandidates{
							RoomStayCandidate: []RoomStayCandidate{
								*roomStayCandidate,
							},
						},
					},
				},
			},
		},
	}
}

func NewOTAHotelPlanDetailRQ(
	// プランの詳細情報取得（合計金額含む）
	planID string,
	hotelCode string,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	TlBookingRoomTypeCode string,
) *OTA_HotelAvailRQ {
	// 日付
	start := util.DateToYYYYMMDD(stayFrom)
	end := util.DateToYYYYMMDD(stayTo)

	// ホテルコード
	hotelRef := HotelRef{HotelCode: hotelCode}

	roomStayCandidate := NewRoomStayCandidate(
		&TlBookingRoomTypeCode,
		nil,
		adult,
		&child,
		roomCount,
		nil, // smokeTypeは指定不可
		nil, // EffectiveDate
		nil, // ExpireDate
	)

	return &OTA_HotelAvailRQ{
		Version:        "1.0",
		PrimaryLangID:  "jpn",
		AvailRatesOnly: util.BoolPtr(true), // プラン情報を返すフラグ
		AvailRequestSegments: AvailRequestSegments{
			AvailRequestSegment: AvailRequestSegment{
				HotelSearchCriteria: HotelSearchCriteria{
					Criterion: Criterion{
						HotelRef: hotelRef,
						StayDateRange: &StayDateRange{
							Start: &start,
							End:   &end,
						},
						RatePlanCandidates: &RatePlanCandidates{
							RatePlanCandidate: []RatePlanCandidate{
								{
									RatePlanCode: &planID,
								},
							},
						},
						RoomStayCandidates: &RoomStayCandidates{
							RoomStayCandidate: []RoomStayCandidate{
								*roomStayCandidate,
							},
						},
					},
				},
			},
		},
	}
}

func NewRoomStayCandidate(
	TlBookingRoomTypeCode *string,
	bedTypeCode *BedTypeCode,
	adult int,
	child *int,
	roomCount int,
	smokeTypes []entity.SmokeType,
	effectiveDate *time.Time,
	expireDate *time.Time,
) *RoomStayCandidate {
	// True：禁煙、False：喫煙、省略：条件指定なし
	var nonSmokingQuery *bool
	if smokeTypes == nil {
		nonSmokingQuery = nil
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeNonSmoking) && entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeSmoking) {
		// 禁煙喫煙両方の場合は条件指定なし
		nonSmokingQuery = nil
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeNonSmoking) {
		nonSmokingQuery = util.BoolPtr(true)
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeSmoking) {
		nonSmokingQuery = util.BoolPtr(false)
	}
	candidate := &RoomStayCandidate{
		Quantity:     &roomCount,
		BedTypeCode:  bedTypeCode,
		RoomTypeCode: TlBookingRoomTypeCode,
		NonSmoking:   nonSmokingQuery,
		GuestCounts: &GuestCounts{
			GuestCount: []GuestCount{
				{
					AgeQualifyingCode: AgeQualifyingAdult,
					Count:             adult,
				},
			},
		},
	}
	if child != nil && *child > 0 {
		candidate.GuestCounts.GuestCount = append(candidate.GuestCounts.GuestCount, GuestCount{
			AgeQualifyingCode: AgeQualifyingChild,
			Count:             *child,
		})
	}
	if effectiveDate != nil {
		effective := util.DateToYYYYMMDD(*effectiveDate)
		candidate.EffectiveDate = &effective
	}
	if expireDate != nil {
		expire := util.DateToYYYYMMDD(*expireDate)
		candidate.ExpireDate = &expire
	}
	return candidate
}

func MealTypeToQuery(mealType entity.MealType) *MealsIncluded {
	var mealMorning *bool = nil
	if mealType.Morning {
		mealMorning = util.BoolPtr(true)
	}
	var mealDinner *bool = nil
	if mealType.Dinner {
		mealDinner = util.BoolPtr(true)
	}
	if mealType.Morning && mealType.Dinner {
		// 朝食と夕食が両方ある場合は条件指定なし
		mealMorning = nil
		mealDinner = nil
	}
	return &MealsIncluded{
		Breakfast: mealMorning,
		Dinner:    mealDinner,
	}
}

func RoomTypeToBedType(rt entity.RoomType) *BedTypeCode {
	var code BedTypeCode
	switch rt {
	case entity.RoomTypeSingle:
		code = BedTypeSingle
	case entity.RoomTypeSemiDouble:
		code = BedTypeSemiDouble
	case entity.RoomTypeTwin:
		code = BedTypeTwin
	case entity.RoomTypeDouble:
		code = BedTypeDouble
	case entity.RoomTypeFourth:
		code = BedTypeFour
	case entity.RoomTypeUnknown:
		return nil
	}
	return &code
}

func BedTypeCodeToRoomType(bt BedTypeCode) entity.RoomType {
	var roomType entity.RoomType
	switch bt {
	case BedTypeSingle:
		roomType = entity.RoomTypeSingle
	case BedTypeDouble:
		roomType = entity.RoomTypeDouble
	case BedTypeTwin:
		roomType = entity.RoomTypeTwin
	case BedTypeFour:
		roomType = entity.RoomTypeFourth
	case BedTypeSemiDouble:
		roomType = entity.RoomTypeSemiDouble
	case BedTypeTatami:
		fallthrough
	case BedTypeTriple:
		fallthrough
	case BedTypeOther:
	default:
		roomType = entity.RoomTypeUnknown
	}
	return roomType
}

func IsNonSmokingToSmokeType(isNonSmoke *bool) entity.SmokeType {
	if isNonSmoke == nil {
		return entity.SmokeTypeUnknown
	}
	if *isNonSmoke {
		return entity.SmokeTypeNonSmoking
	} else {
		return entity.SmokeTypeSmoking
	}
}
