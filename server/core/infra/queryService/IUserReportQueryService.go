package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserReportQueryService interface {
	GetByID(id uuid.UUID) (*entity.UserReport, error)
	GetAll(pager *types.PageQuery) ([]*entity.UserReport, error)
}
