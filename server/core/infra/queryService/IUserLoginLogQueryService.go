package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserLoginLogQueryService interface {
	GetList(userID uuid.UUID, pager *types.PageQuery) ([]*entity.UserLoginLog, *types.PageResponse, error)
}
