package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
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
	adult uint8,
	child uint8,
	roomType []entity.RoomType,
	pager *types.PageQuery,
) ([]*entity.Plan, *errors.DomainError) {

	plans, err := u.planQuery.Search(
		stayStore,
		stayFrom,
		stayTo,
		adult,
		child,
		roomType,
		pager,
	)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return plans, nil
}
