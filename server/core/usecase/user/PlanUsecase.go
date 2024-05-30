package user

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"

	"github.com/google/uuid"
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
	stayStoreIds []uuid.UUID,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes *[]entity.SmokeType,
	mealType *entity.MealType,
	roomTypes *[]entity.RoomType,
) (*[]entity.PlanCandidate, *errors.DomainError) {
	var stayStores []*entity.StayableStore
	var err error

	if len(stayStoreIds) == 0 {
		stayStores, err = u.storeQuery.GetStayables()
		if err != nil {
			return nil, errors.NewDomainError(errors.QueryError, err.Error())
		}
	} else {
		for _, storeID := range stayStoreIds {
			stayStore, err := u.storeQuery.GetStayableByID(storeID)
			if err != nil {
				return nil, errors.NewDomainError(errors.QueryError, err.Error())
			}
			if stayStore == nil {
				return nil, errors.NewDomainError(errors.InvalidParameter, "指定された宿泊施設が見つかりませんでした。")
			}

			stayStores = append(stayStores, stayStore)
		}
	}

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

	plans, err := u.planQuery.Search(
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
	var candidates []entity.PlanCandidate
	nights := stayTo.Sub(stayFrom).Hours() / 24
	guestCount := adult + child
	for _, plan := range *plans {
		candidate := entity.NewPlanCandidate(&plan, int(nights), guestCount)
		candidates = append(candidates, *candidate)
	}
	return &candidates, nil
}

func (u *PlanUsecase) GetDetail(
	PrinID string,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	roomTypes *entity.RoomType,
	stayStoreID uuid.UUID,
) (*entity.Plan, *errors.DomainError) {
	stayStore, err := u.storeQuery.GetStayableByID(stayStoreID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if stayStore == nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, "指定された宿泊施設が見つかりませんでした。")
	}
	plan, err := u.planQuery.GetPlanDetailByID(PrinID, stayStore, stayFrom, stayTo, adult, child, roomCount, *roomTypes)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return plan, nil
}
