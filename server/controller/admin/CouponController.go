package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/api/v1/shared"
	"server/controller"
	usecase "server/core/usecase/admin"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminCouponController struct {
	couponUseCase usecase.AdminCouponUsecase
}

var _ adminv1connect.AdminCouponControllerClient = &AdminCouponController{}

func NewAdminCouponController(couponUsecase *usecase.AdminCouponUsecase) *AdminCouponController {
	return &AdminCouponController{
		couponUseCase: *couponUsecase,
	}
}

func (ac *AdminCouponController) GetUserCouponList(ctx context.Context, req *connect.Request[admin.UserIDRequest]) (*connect.Response[admin.UserAttachedCouponsResponse], error) {
	userID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	entities, domaiErr := ac.couponUseCase.GetUsersCouponList(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var attachedCoupons []*shared.UserAttachedCoupon
	for _, entity := range entities {
		cpn := &shared.Coupon{
			ID:                entity.ID.String(),
			Name:              entity.Coupon.Name,
			CouponType:        shared.CouponType(entity.Coupon.CouponType.ToInt()),
			DiscountAmount:    uint32(entity.Coupon.DiscountAmount),
			ExpireAt:          timestamppb.New(entity.Coupon.ExpireAt),
			IsCombinationable: entity.Coupon.IsCombinationable,
		}
		atcCpn := &shared.UserAttachedCoupon{
			UserID: userID.String(),
			Coupon: cpn,
			UsedAt: timestamppb.New(entity.ExpireAt),
		}
		attachedCoupons = append(attachedCoupons, atcCpn)
	}
	result := &admin.UserAttachedCouponsResponse{
		UserAttachedCoupons: attachedCoupons,
	}

	return connect.NewResponse(result), nil
}

func (ac *AdminCouponController) CreateCustomCoupon(ctx context.Context, req *connect.Request[admin.CreateCustomCouponRequest]) (*connect.Response[emptypb.Empty], error) {
	_, domaiErr := ac.couponUseCase.CreateCustomCoupon(
		req.Msg.Name,
		uint(req.Msg.DiscountAmount),
		req.Msg.ExpireAt.AsTime(),
		req.Msg.IsCombinationable,
		req.Msg.Notices,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AdminCouponController) SaveCustomCoupon(ctx context.Context, req *connect.Request[admin.CouponIDRequest]) (*connect.Response[emptypb.Empty], error) {
	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	domaiErr := ac.couponUseCase.SaveCustomCoupon(couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AdminCouponController) AttachCustomCouponToAllUser(ctx context.Context, req *connect.Request[admin.CouponIDRequest]) (*connect.Response[admin.AffectedCountResponse], error) {
	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	count, domaiErr := ac.couponUseCase.AttachCustomCouponToAllUser(couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	result := &admin.AffectedCountResponse{
		AffectedUserCount: uint32(*count),
	}
	return connect.NewResponse(result), nil
}
