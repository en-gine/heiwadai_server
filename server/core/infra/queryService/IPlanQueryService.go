package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"
	"time"
)

type IPlanQueryService interface {
	Search(
		stayStore []entity.Store,
		stayFrom time.Time,
		stayTo time.Time,
		adult uint8,
		child uint8,
		roomType []entity.RoomType,
		pager *types.PageQuery) ([]*entity.Plan, error)
}
