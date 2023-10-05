package booking

import (
	"fmt"
	"reflect"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/avail"
	"server/infrastructure/booking/util"

	uuid "github.com/google/uuid"
)

var _ queryservice.IPlanQueryService = &PlanQuery{}

type PlanQuery struct{}

func NewPlanQuery() *PlanQuery {
	return &PlanQuery{}
}

func (p *PlanQuery) GetMyBooking(userID uuid.UUID) ([]*entity.Plan, error) {
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
) ([]*entity.Plan, error) {
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

	_, err := Request[avail.OTAHotelAvailRQ, avail.OTAHotelAvailRS](reqBody)
	if err != nil {
		return nil, err
	}
	return nil, nil
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

	if reflect.DeepEqual(smokeTypes, []entity.SmokeType{entity.SmokeTypeNonSmoking}) {
		nonSmokingQuery = util.BoolPtr(true)
	}

	if reflect.DeepEqual(smokeTypes, []entity.SmokeType{entity.SmokeTypeSmoking}) {
		nonSmokingQuery = util.BoolPtr(false)
	}

	if reflect.DeepEqual(smokeTypes, []entity.SmokeType{entity.SmokeTypeNonSmoking, entity.SmokeTypeSmoking}) {
		nonSmokingQuery = nil
	}

	fmt.Print(nonSmokingQuery)
	nonSmokingQuery = util.BoolPtr(true)

	// 部屋プラン検索
	var roomStayCandidates []avail.RoomStayCandidate
	var bedTypeCode *avail.BedTypeCode = nil

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
		HotelStayOnly: util.BoolPtr(true), // ホテル情報のみを返すフラグ。不明
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
	}
	return &code
}
