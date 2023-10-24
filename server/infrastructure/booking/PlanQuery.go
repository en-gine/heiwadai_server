package booking

import (
	"errors"
	"strconv"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/avail"
	"server/infrastructure/booking/util"
	"server/infrastructure/logger"

	uuid "github.com/google/uuid"
)

var _ queryservice.IPlanQueryService = &PlanQuery{}

type PlanQuery struct {
	storeQuery queryservice.IStoreQueryService
}

func NewPlanQuery(storeQuery queryservice.IStoreQueryService) *PlanQuery {
	return &PlanQuery{
		storeQuery: storeQuery,
	}
}

func (p *PlanQuery) GetMyBooking(userID uuid.UUID) (*[]entity.Plan, error) {
	return nil, nil
}

func (p *PlanQuery) Search(
	stayStore []entity.Store,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes *[]entity.SmokeType,
	mealType *entity.MealType,
	roomTypes *[]entity.RoomType,
) (*[]entity.Plan, error) {
	reqBody := NewOTAHotelAvailRQ(
		[]string{"E69502"},
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		smokeTypes,
		mealType,
		roomTypes,
	)

	res, err := Request[avail.OTAHotelAvailRQ, avail.OTAHotelAvailRS](reqBody)
	if err != nil {
		return nil, err
	}

	plans, err := p.AvailRSToPlans(res)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (p *PlanQuery) AvailRSToPlans(res *avail.OTAHotelAvailRS) (*[]entity.Plan, error) {
	var plans []entity.Plan

	if res.RoomStays == nil {
		logger.Errorf("AvailRSToPlans Return Error: %v", res)
		return nil, errors.New("RoomStays is nil")
	}
	for _, roomStay := range res.RoomStays.RoomStay {
		hotelCode := roomStay.RPH
		if hotelCode == nil {
			logger.Errorf("AvailRSToPlans Return Error: %v", res)
			return nil, errors.New("hotelCode is nil")
		}
		stayable, err := p.storeQuery.GetStayableByBookingID(*hotelCode)
		if err != nil {
			return nil, err
		}
		roomType := roomStay.RoomTypes.RoomType[0]
		room := BedTypeCodeToRoomType(roomType.BedTypeCode)
		smokeType := IsNonSmokingToSmokeType(roomType.NonSmoking) // true false nil

		// roomName := roomType.RoomDescription.Name
		// roomText := roomType.RoomDescription.Text
		// roomImageUrl := roomType.RoomDescription.URL

		for _, plan := range *roomStay.RatePlans.RatePlan {
			planID := plan.RatePlanCode
			planName := plan.RatePlanName

			var planImageURL string = ""
			if len(*plan.RatePlanDescription.URL) > 0 {
				planImageURL = *(*plan.RatePlanDescription.URL)[0].Value
			}

			var planOverView string = ""
			if len(*plan.RatePlanDescription.Text) > 0 {
				planOverView = *(*plan.RatePlanDescription.Text)[0].Value
			}

			var planPrice uint64
			planPrice, _ = strconv.ParseUint(*plan.RatePlanType, 10, 64)

			var IncludeBreakfast bool = *plan.MealsIncluded.Breakfast
			var IncludeDinner bool = *plan.MealsIncluded.Dinner

			plan := entity.Plan{
				ID:       *planID,
				Title:    *planName,
				Price:    uint(planPrice),
				ImageURL: planImageURL,
				RoomType: room,
				MealType: entity.MealType{
					Morning: IncludeBreakfast,
					Dinner:  IncludeDinner,
				},
				SmokeType: smokeType,
				OverView:  planOverView,
				Store:     *stayable,
			}

			plans = append(plans, plan)
		}

	}

	return &plans, nil
}

func NewOTAHotelAvailRQ(
	hotelCodes []string,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes *[]entity.SmokeType,
	mealType *entity.MealType,
	roomTypes *[]entity.RoomType,
) *avail.OTAHotelAvailRQ {
	// 日付
	start := util.DateToYYYYMMDD(stayFrom)
	end := util.DateToYYYYMMDD(stayTo)

	// ホテルコード
	var hotelRef []avail.HotelRef
	for _, hotelCode := range hotelCodes {
		hotelRef = append(hotelRef, avail.HotelRef{HotelCode: hotelCode})
	}
	// True：禁煙、False：喫煙、省略：条件指定なし
	var nonSmokingQuery *bool
	if smokeTypes == nil {
		nonSmokingQuery = nil
	} else if entity.IncludeSmokeType(*smokeTypes, entity.SmokeTypeNonSmoking) {
		nonSmokingQuery = util.BoolPtr(true)
	} else if entity.IncludeSmokeType(*smokeTypes, entity.SmokeTypeSmoking) {
		nonSmokingQuery = util.BoolPtr(false)
	} else if entity.IncludeSmokeType(*smokeTypes, entity.SmokeTypeNonSmoking) && entity.IncludeSmokeType(*smokeTypes, entity.SmokeTypeSmoking) {
		// 禁煙喫煙両方の場合は条件指定なし
		nonSmokingQuery = nil
	}

	// 部屋プラン検索
	var roomStayCandidates []avail.RoomStayCandidate
	var bedTypeCode *avail.BedTypeCode

	var mealMorning *bool
	var mealDinner *bool
	if mealType != nil {
		mealMorning = &mealType.Morning
		mealDinner = &mealType.Dinner
	} else {
		mealMorning = nil
		mealDinner = nil
	}

	if roomTypes != nil && len(*roomTypes) > 0 {
		for _, rt := range *roomTypes {
			bedTypeCode = RoomTypeToBedType(&rt)
			candidate := avail.RoomStayCandidate{
				Quantity:    &roomCount,
				BedTypeCode: bedTypeCode,
				NonSmoking:  nonSmokingQuery,
				GuestCounts: &avail.GuestCounts{
					GuestCount: []avail.GuestCount{
						{
							AgeQualifyingCode: avail.AgeQualifyingAdult,
							Count:             adult,
						},
					},
				},
			}
			roomStayCandidates = append(roomStayCandidates, candidate)
		}
	} else {
		candidate := avail.RoomStayCandidate{
			Quantity:    &roomCount,
			NonSmoking:  nonSmokingQuery,
			BedTypeCode: nil,
			GuestCounts: &avail.GuestCounts{
				GuestCount: []avail.GuestCount{
					{
						AgeQualifyingCode: avail.AgeQualifyingAdult,
						Count:             adult,
					},
					{
						AgeQualifyingCode: avail.AgeQualifyingChild,
						Count:             child,
					},
				},
			},
		}
		roomStayCandidates = append(roomStayCandidates, candidate)
	}

	return &avail.OTAHotelAvailRQ{
		Version:       "1.0",
		PrimaryLangID: "jpn",
		HotelStayOnly: util.BoolPtr(true), // ホテル情報のみを返すフラグ。trueにしないと返ってこない
		PricingMethod: avail.PricingMethodLowestperstay,
		AvailRequestSegments: avail.AvailRequestSegments{
			AvailRequestSegment: avail.AvailRequestSegment{
				HotelSearchCriteria: avail.HotelSearchCriteria{
					Criterion: avail.Criterion{
						HotelRef: hotelRef,
						RatePlanCandidates: &avail.RatePlanCandidates{ // 食事タイプ
							RatePlanCandidate: []avail.RatePlanCandidate{
								{
									MealsIncluded: &avail.MealsIncluded{
										Breakfast: mealMorning,
										Dinner:    mealDinner,
									},
								},
							},
						},
						StayDateRange: &avail.StayDateRange{
							Start: &start,
							End:   &end,
						},
						RoomStayCandidates: &avail.RoomStayCandidates{
							RoomStayCandidate: roomStayCandidates,
						},
					},
				},
			},
		},
	}
}

func RoomTypeToBedType(rt *entity.RoomType) *avail.BedTypeCode {
	var code avail.BedTypeCode
	if rt == nil {
		return nil
	}
	switch *rt {
	case entity.RoomTypeSingle:
		code = avail.BedTypeSingle
	case entity.RoomTypeSemiDouble:
		code = avail.BedTypeDouble
	case entity.RoomTypeTwin:
		code = avail.BedTypeTwin
	case entity.RoomTypeDouble:
		code = avail.BedTypeDouble
	case entity.RoomTypeFourth:
		code = avail.BedTypeFour
	case entity.RoomTypeUnknown:
		return nil
	}
	return &code
}

func BedTypeCodeToRoomType(bt *avail.BedTypeCode) entity.RoomType {
	var roomType entity.RoomType
	switch *bt {
	case avail.BedTypeSingle:
		roomType = entity.RoomTypeSingle
	case avail.BedTypeDouble:
		roomType = entity.RoomTypeDouble
	case avail.BedTypeTwin:
		roomType = entity.RoomTypeTwin
	case avail.BedTypeFour:
		roomType = entity.RoomTypeFourth
	case avail.BedTypeTatami:
		fallthrough
	case avail.BedTypeSemiDouble:
		fallthrough
	case avail.BedTypeTriple:
		fallthrough
	case avail.BedTypeOther:
	default:
		roomType = entity.RoomTypeUnknown
	}
	return roomType
}

func IsNonSmokingToSmokeType(isNonSmoke *bool) entity.SmokeType {
	switch *isNonSmoke {
	case true:
		return entity.SmokeTypeNonSmoking
	case false:
		return entity.SmokeTypeNonSmoking
	}
	return entity.SmokeTypeUnknown
}
