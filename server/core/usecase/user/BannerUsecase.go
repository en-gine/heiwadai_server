package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
)

type BannerUsecase struct {
	bannerQuery queryservice.IBannerQueryService
}

func NewBannerUsecase(bannerQuery queryservice.IBannerQueryService) *BannerUsecase {
	return &BannerUsecase{
		bannerQuery: bannerQuery,
	}
}

func (u *BannerUsecase) GetList() ([]*entity.Banner, *errors.DomainError) {
	banners, err := u.bannerQuery.GetAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return banners, nil
}
