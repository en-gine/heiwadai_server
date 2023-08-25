package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"time"
)

type PlanUsecase struct {
	planQuery queryservice.IPlanQueryService
}

func NewPlanUsecase(planQuery queryservice.IPlanQueryService) *PlanUsecase {
	return &PlanUsecase{
		planQuery: planQuery,
	}
}

func (u *PlanUsecase) Search(
	stayStore []entity.Store,
	stayFrom time.Time,
	stayTo time.Time,
	adult int,
	child int,
	roomCount int,
	smokeTypes []entity.SmokeType,
	mealType entity.MealType,
	roomTypes []entity.RoomType,
) ([]*entity.Plan, *errors.DomainError) {

	plans, err := u.planQuery.Search(
		stayStore,
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
