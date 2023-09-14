package user

import (
	"context"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BannerController struct {
	bannerUseCase usecase.BannerUsecase
}

var _ userv1connect.BannerControllerClient = &BannerController{}

func NewBannerController(bannerUsecase *usecase.BannerUsecase) *BannerController {
	return &BannerController{
		bannerUseCase: *bannerUsecase,
	}
}

func (ac *BannerController) GetBanner(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.BannerResponse], error) {
	banners, domaiErr := ac.bannerUseCase.GetList()
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var resBanners []*user.Banner
	for _, banner := range banners {
		resBanners = append(resBanners, &user.Banner{
			ImageURL: banner.ImageURL,
			URL:      banner.URL,
		})
	}

	return connect.NewResponse(&user.BannerResponse{
		Banners: resBanners,
	}), nil
}
