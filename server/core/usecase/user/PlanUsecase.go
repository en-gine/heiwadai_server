package user

import (
	"reflect"
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
)

type PlanUsecase struct {
	planQuery  queryservice.IPlanQueryService
	storeQuery queryservice.IStoreQueryService
}

func NewPlanUsecase(planQuery queryservice.IPlanQueryService, storeQuery queryservice.IStoreQueryService) *PlanUsecase {
	return &PlanUsecase{
		planQuery:  planQuery,
		storeQuery: storeQuery,
	}
}

func (u *PlanUsecase) Search(
	stayStores []*entity.StayableStore,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes *[]entity.SmokeType,
	mealType *entity.MealType,
	roomTypes *[]entity.RoomType,
) (*[]entity.PlanCandidate, *errors.DomainError) {
	var err error

	if roomTypes == nil || len(*roomTypes) == 0 {
		roomTypes = &entity.RoomTypeAll
	}

	if smokeTypes == nil || len(*smokeTypes) == 0 {
		smokeTypes = &entity.SmokeTypeAll
	}

	if mealType == nil {
		mealType = &entity.MealType{
			Morning: true,
			Dinner:  true,
		}
	}

	candidates, err := u.planQuery.Search(
		stayStores,
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		*smokeTypes,
		*mealType,
		*roomTypes,
	)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return candidates, nil
}

func (u *PlanUsecase) GetDetail(
	PrinID string,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	TlBookingRoomTypeCode string, // 検索時に返ってくるTLBookingの部屋タイプコード
	stayStore *entity.StayableStore,
) (*entity.PlanStayDetail, *errors.DomainError) {

	plan, err := u.planQuery.GetPlanDetailByID(PrinID, stayStore, stayFrom, stayTo, adult, child, roomCount, TlBookingRoomTypeCode)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if plan == nil || plan.Plan == nil || reflect.DeepEqual(plan, entity.PlanStayDetail{}) || plan.Plan.ID == "" {
		return nil, errors.NewDomainError(errors.CancelButNeedFeedBack, "プランの詳細が取得できませんでした。\n既に販売終了している可能性があります。")
	}
	return plan, nil
}
