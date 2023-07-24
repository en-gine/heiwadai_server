package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IBannerQueryService interface {
	GetByID(id uuid.UUID) (*entity.Banner, error)
	GetActiveAll(pager *types.PageQuery) ([]*entity.Banner, error)
	GetAll(pager *types.PageQuery) ([]*entity.Banner, error) //statusは不問
}
