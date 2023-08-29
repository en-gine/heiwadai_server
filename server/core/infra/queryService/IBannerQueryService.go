package queryservice

import (
	"server/core/entity"
)

type IBannerQueryService interface {
	GetByID(id int) (*entity.Banner, error)
	GetAll() ([]*entity.Banner, error) //statusは不問
}
