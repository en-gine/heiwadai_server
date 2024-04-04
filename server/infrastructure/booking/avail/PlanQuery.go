package avail

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/util"
	"server/infrastructure/env"
	"server/infrastructure/logger"

	uuid "github.com/google/uuid"
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

func (p *PlanQuery) GetCalendar(
	planID string, //nolint:all
	storeID uuid.UUID,
	night int,
	adult int,
	child int,
	roomCount int,
	smokeTypes []entity.SmokeType,
	mealType entity.MealType,
	roomType entity.RoomType,
	fromDate time.Time,
	toDate time.Time,
) (*entity.PlanCalendar, error) {
	store, err := p.storeQuery.GetStayableByID(storeID)
	if err != nil {
		return nil, err
	}
	hotelCode := store.BookingSystemID

	if env.GetEnv(env.TlbookingIsTest) == "true" {
		hotelCode = "E69502"
		planID = "14030824" //nolint:all
	}

	mealsIncluded := MealTypeToQuery(mealType)

	roomStayCandidate := NewRoomStayCandidate(
		roomType,
		adult,
		&child,
		roomCount,
		smokeTypes,
		&fromDate,
		&toDate,
	)

	nightFormat := fmt.Sprintf("P%dN", night)
	reqBody := &OTA_HotelAvailRQ{
		Version:        "1.0",
		PrimaryLangID:  "jpn",
		AvailRatesOnly: util.BoolPtr(true),
		AvailRequestSegments: AvailRequestSegments{
			AvailRequestSegment: AvailRequestSegment{
				HotelSearchCriteria: HotelSearchCriteria{
					Criterion: Criterion{
						HotelRef: []HotelRef{
							{HotelCode: hotelCode},
						},
						RatePlanCandidates: &RatePlanCandidates{ // 食事タイプ
							RatePlanCandidate: []RatePlanCandidate{
								{
									MealsIncluded: mealsIncluded,
								},
							},
						},
						StayDateRange: &StayDateRange{
							Duration: &nightFormat,
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

	request := NewEnvelopeRQ(TLBookingUser, TLBookingPass, reqBody)

	res, err := util.Request[EnvelopeRQ, EnvelopeRS](TLBookingSearchURL, request)
	if err != nil {
		return nil, err
	}

	if res.Body.OTA_HotelAvailRS.Errors != nil {
		errs := res.Body.OTA_HotelAvailRS.Errors
		msg := errs.Error[0].ShortText
		return nil, errors.New(msg)
	}

	plans, err := p.AvailRSToPlans(res)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	plan := (*plans)[0]
	statases, err := p.AvailRSToCalendarStatus(res)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &entity.PlanCalendar{
		Plan:       plan,
		DateStatus: *statases,
	}, nil
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
) (*[]entity.Plan, error) {
	bookingIDs := []string{}
	for _, store := range stores {
		bookingIDs = append(bookingIDs, store.BookingSystemID)
	}

	if env.GetEnv(env.TlbookingIsTest) == "true" {
		bookingIDs = []string{"E69502"}
	}

	type result struct {
		plans *[]entity.Plan
		err   error
	}

	resultsCh := make(chan result, len(roomTypes))

	for _, roomType := range roomTypes {
		go func(rt entity.RoomType) {
			reqBody := NewOTAHotelAvailRQ(
				bookingIDs,
				stayFrom,
				stayTo,
				adult,
				child,
				roomCount,
				smokeTypes,
				mealType,
				rt,
			)
			request := NewEnvelopeRQ(TLBookingUser, TLBookingPass, reqBody)
			res, err := util.Request[EnvelopeRQ, EnvelopeRS](TLBookingSearchURL, request)
			if err != nil {
				resultsCh <- result{nil, err}
				return
			}

			if res.Body.OTA_HotelAvailRS.Errors != nil {
				errs := res.Body.OTA_HotelAvailRS.Errors
				msg := errs.Error[0].ShortText
				logger.Error(msg)
				resultsCh <- result{nil, errors.New(msg)}
				return
			}

			plans, err := p.AvailRSToPlans(res)
			if err != nil {
				logger.Error(err.Error())
				resultsCh <- result{nil, err}
				return
			}
			resultsCh <- result{plans, nil}
		}(roomType)
	}

	var allPlans []entity.Plan
	for i := 0; i < len(roomTypes); i++ {
		res := <-resultsCh
		if res.err != nil {
			return nil, res.err
		}
		allPlans = append(allPlans, *res.plans...)
	}
	return &allPlans, nil
}

func (p *PlanQuery) AvailRSToPlans(res *EnvelopeRS) (*[]entity.Plan, error) {
	var plans []entity.Plan
	body := res.Body.OTA_HotelAvailRS
	for _, roomStay := range body.RoomStays.RoomStay {
		hotelCode := roomStay.RPH
		var stayable *entity.StayableStore
		var err error
		if env.GetEnv(env.TlbookingIsTest) != "true" {
			stayable, err = p.storeQuery.GetStayableByBookingID(hotelCode)
		} else {
			stayable = &entity.StayableStore{}
		}
		if err != nil {
			return nil, err
		}
		if len(roomStay.RoomTypes.RoomType) == 0 { // room not found
			return &[]entity.Plan{}, nil
		}

		roomType := roomStay.RoomTypes.RoomType[0]
		room := BedTypeCodeToRoomType(roomType.BedTypeCode)
		smokeType := IsNonSmokingToSmokeType(roomType.NonSmoking) // true false nil

		// roomName := roomType.RoomDescription.Name
		// roomText := roomType.RoomDescription.Text
		// roomImageUrl := roomType.RoomDescription.URL
		amount := roomStay.RoomRates.RoomRate[0].Total.AmountAfterTax
		var planPrice uint64
		planPrice, _ = strconv.ParseUint(amount, 10, 64)

		for _, plan := range roomStay.RatePlans.RatePlan {
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

			plan := entity.Plan{
				ID:       planID,
				Title:    planName,
				Price:    uint(planPrice),
				ImageURL: planImageURL,
				RoomType: room,
				MealType: entity.MealType{
					Morning: IncludeBreakfast,
					Dinner:  IncludeDinner,
				},
				SmokeType: smokeType,
				OverView:  planOverView,
				StoreID:   stayable.ID,
			}

			plans = append(plans, plan)
		}

	}

	return &plans, nil
}

func (p *PlanQuery) AvailRSToCalendarStatus(res *EnvelopeRS) (*[]entity.DateStatus, error) {
	var dateStatuses []entity.DateStatus
	body := res.Body.OTA_HotelAvailRS
	for _, roomStay := range body.RoomStays.RoomStay {
		for _, roomRate := range roomStay.RoomRates.RoomRate {
			amount := roomRate.Total.AmountAfterTax
			var planPrice uint64
			planPrice, _ = strconv.ParseUint(amount, 10, 64)
			date, err := util.YYYYMMDDToDate(roomRate.EffectiveDate)
			if err != nil {
				return nil, err
			}
			dateStatus := entity.DateStatus{
				Date:  date,
				Price: uint(planPrice),
			}
			dateStatuses = append(dateStatuses, dateStatus)
		}
	}
	return &dateStatuses, nil
}

func NewOTAHotelAvailRQ(
	hotelCodes []string,
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
	var hotelRef []HotelRef
	for _, hotelCode := range hotelCodes {
		hotelRef = append(hotelRef, HotelRef{HotelCode: hotelCode})
	}

	mealsIncluded := MealTypeToQuery(mealType)

	roomStayCandidate := NewRoomStayCandidate(
		roomType,
		adult,
		&child,
		roomCount,
		smokeTypes,
		nil, // EffectiveDate
		nil, // ExpireDate
	)

	return &OTA_HotelAvailRQ{
		Version:       "1.0",
		PrimaryLangID: "jpn",
		HotelStayOnly: util.BoolPtr(true), // ホテル情報のみを返すフラグ。trueにしないと返ってこない
		PricingMethod: PricingMethodLowestperstay,
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

func NewRoomStayCandidate(
	roomType entity.RoomType,
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
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeNonSmoking) {
		nonSmokingQuery = util.BoolPtr(true)
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeSmoking) {
		nonSmokingQuery = util.BoolPtr(false)
	} else if entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeNonSmoking) && entity.IncludeSmokeType(smokeTypes, entity.SmokeTypeSmoking) {
		// 禁煙喫煙両方の場合は条件指定なし
		nonSmokingQuery = nil
	}

	bedTypeCode := RoomTypeToBedType(roomType)
	candidate := &RoomStayCandidate{
		Quantity:    &roomCount,
		BedTypeCode: bedTypeCode,
		NonSmoking:  nonSmokingQuery,
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
	var mealMorning *bool
	var mealDinner *bool
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
		code = BedTypeDouble
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
	case BedTypeTatami:
		fallthrough
	case BedTypeSemiDouble:
		fallthrough
	case BedTypeTriple:
		fallthrough
	case BedTypeOther:
	default:
		roomType = entity.RoomTypeUnknown
	}
	return roomType
}

func IsNonSmokingToSmokeType(isNonSmoke bool) entity.SmokeType {
	switch isNonSmoke {
	case true:
		return entity.SmokeTypeNonSmoking
	case false:
		return entity.SmokeTypeNonSmoking
	}
	return entity.SmokeTypeUnknown
}
