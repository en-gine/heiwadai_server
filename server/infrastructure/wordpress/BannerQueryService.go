package wordpress

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
)

var _ queryservice.IBannerQueryService = &BannerQueryService{}

type BannerQueryService struct {
}

func NewBannerQueryService() *BannerQueryService {

	return &BannerQueryService{}
}

func (pq *BannerQueryService) GetByID(id int) (*entity.Banner, error) {

	return &entity.Banner{}, nil
}

func (pq *BannerQueryService) GetAll() ([]*entity.Banner, error) {
	return nil, nil

}
