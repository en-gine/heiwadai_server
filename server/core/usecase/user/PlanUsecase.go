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
	smokeTypes []entity.SmokeType,
	mealType entity.MealType,
	roomTypes []entity.RoomType,
) (*[]entity.Plan, *errors.DomainError) {
	var stayStores []*entity.StayableStore
	var err error

	if stayStoreIds == nil || len(stayStoreIds) == 0 {
		stayStores, err = u.storeQuery.GetStayables()
		if err != nil {
			return nil, errors.NewDomainError(errors.QueryError, err.Error())
		}
	} else {
		for _, storeId := range stayStoreIds {
			stayStore, err := u.storeQuery.GetStayableByID(storeId)
			if err != nil {
				return nil, errors.NewDomainError(errors.QueryError, err.Error())
			}
			if stayStore == nil {
				return nil, errors.NewDomainError(errors.InvalidParameter, "指定された宿泊施設が見つかりませんでした。")
			}

			stayStores = append(stayStores, stayStore)
		}
	}

	plans, err := u.planQuery.Search(
		stayStores,
		stayFrom,
		stayTo,
		adult,
		child,
		roomCount,
		&smokeTypes,
		&mealType,
		&roomTypes,
	)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return plans, nil
}
