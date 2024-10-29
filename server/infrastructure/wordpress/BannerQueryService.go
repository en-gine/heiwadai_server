package wordpress

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/logger"
	"server/infrastructure/wordpress/api"
	"server/infrastructure/wordpress/types"
)

var _ queryservice.IBannerQueryService = &BannerQueryService{}

type BannerQueryService struct {
}

func NewBannerQueryService() *BannerQueryService {

	return &BannerQueryService{}
}

func (pq *BannerQueryService) GetAll() ([]*entity.Banner, error) {
	banners, err := api.GetWPBanners()
	if err != nil {
		logger.Errorf("Error: %v\n", err)
		return nil, err
	}
	var entities []*entity.Banner
	if banners == nil {
		return entities, nil
	}
	for _, wpbanner := range *banners {
		entities = append(entities, WPBannerToEntity(&wpbanner))
	}
	return entities, nil
}

func WPBannerToEntity(wpbanner *types.WPBanner) (entitie *entity.Banner) {
	entity := &entity.Banner{
		ImageURL: wpbanner.IMAGE,
		URL:      wpbanner.URL,
	}
	return entity
}
