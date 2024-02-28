package admin

import (
	"context"
	"time"

	"server/api/v1/cron"
	cronv1connect "server/api/v1/cron/cronconnect"
	"server/controller"
	usecase "server/core/usecase/cron"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CronCouponController struct {
	couponUseCase usecase.CronCouponUsecase
}

var _ cronv1connect.CronCouponControllerClient = &CronCouponController{}

func NewCronCouponController(couponUsecase *usecase.CronCouponUsecase) *CronCouponController {
	return &CronCouponController{
		couponUseCase: *couponUsecase,
	}
}

func (ac *CronCouponController) BulkIssueBirthdayCoupon(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[cron.AffectedCountResponse], error) {
	count, domaiErr := ac.couponUseCase.BulkAttachBirthdayCoupon(int(time.Now().Month()))
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	result := &cron.AffectedCountResponse{
		AffectedUserCount: uint32(*count),
	}
	return connect.NewResponse(result), nil
}
