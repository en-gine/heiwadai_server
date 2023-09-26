package booking

import (
	"reflect"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/avail"
	"server/infrastructure/booking/util"
	"time"

	uuid "github.com/google/uuid"
)

var _ queryservice.IPlanQueryService = &PlanQuery{}

type PlanQuery struct {
}

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
	roomTypes *[]entity.RoomType) ([]*entity.Plan, error) {

	reqBody := NewOTAHotelAvailRQ(
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

	//True：禁煙、False：喫煙、省略：条件指定なし
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

	// 部屋プラン検索
	var roomStayCandidates []avail.RoomStayCandidate
	var bedTypeCode avail.BedTypeCode

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
			switch rt {
			case entity.RoomTypeSingle:
				bedTypeCode = avail.BedTypeSingle
			case entity.RoomTypeSemiDouble:
				bedTypeCode = avail.BedTypeDouble
			case entity.RoomTypeTwin:
				bedTypeCode = avail.BedTypeTwin
			case entity.RoomTypeDouble:
				bedTypeCode = avail.BedTypeDouble
			case entity.RoomTypeFourth:
				bedTypeCode = avail.BedTypeFour
			}
			candidate := avail.RoomStayCandidate{
				Quantity:    &roomCount,
				NonSmoking:  nonSmokingQuery,
				BedTypeCode: &bedTypeCode,
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
	} else {
		candidate := avail.RoomStayCandidate{
			Quantity:    &roomCount,
			NonSmoking:  nonSmokingQuery,
			BedTypeCode: &bedTypeCode,
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
		SummaryOnly:   false,
		AvailRequestSegments: avail.AvailRequestSegments{
			AvailRequestSegment: avail.AvailRequestSegment{
				HotelSearchCriteria: avail.HotelSearchCriteria{
					Criterion: avail.Criterion{
						HotelRef: []avail.HotelRef{
							{
								HotelCode: "HND0001",
							},
						},
						RatePlanCandidates: &avail.RatePlanCandidates{ //食事タイプ
							RatePlanCandidate: []avail.RatePlanCandidate{
								{
									MealsIncluded: avail.MealsIncluded{
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
