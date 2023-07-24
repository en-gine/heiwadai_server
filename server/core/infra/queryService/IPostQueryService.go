package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IPostQueryService interface {
	GetByID(id uuid.UUID) (*entity.Post, error)
	GetActiveAll(pager *types.PageQuery) ([]*entity.Post, error)
	GetAll(pager *types.PageQuery) ([]*entity.Post, error) //statusは不問
}
