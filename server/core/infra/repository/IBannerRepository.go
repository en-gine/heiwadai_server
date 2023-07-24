package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IBannerRepository interface {
	Save(updateBanner *entity.Banner) error
	Delete(bannerID uuid.UUID) error
}
