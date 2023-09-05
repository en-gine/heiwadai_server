package queryservice

import (
	"server/core/entity"
)

type IBannerQueryService interface {
	GetAll() ([]*entity.Banner, error)
}
