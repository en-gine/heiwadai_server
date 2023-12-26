package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/api/v1/shared"
	"server/controller"
	"server/core/entity"
	"server/core/infra/queryService/types"
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
		cpn := EntityToResCoupon(entity.Coupon)

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

func (ac *AdminCouponController) CreateCustomCoupon(ctx context.Context, req *connect.Request[admin.CreateCustomCouponRequest]) (*connect.Response[shared.Coupon], error) {
	entity, domaiErr := ac.couponUseCase.CreateCustomCoupon(
		req.Msg.Name,
		uint(req.Msg.DiscountAmount),
		req.Msg.ExpireAt.AsTime(),
		req.Msg.IsCombinationable,
		req.Msg.Notices,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	cpn := EntityToResCoupon(entity)

	return connect.NewResponse(cpn), nil
}

func (ac *AdminCouponController) GetCustomCouponByID(ctx context.Context, req *connect.Request[admin.CouponIDRequest]) (*connect.Response[shared.Coupon], error) {
	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	entity, domaiErr := ac.couponUseCase.GetCustomCouponByID(couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	cpn := EntityToResCoupon(entity)

	return connect.NewResponse(cpn), nil
}

func (ac *AdminCouponController) GetCustomCouponList(ctx context.Context, req *connect.Request[shared.Pager]) (*connect.Response[admin.CouponListResponse], error) {
	perPage := int(*req.Msg.PerPage)
	page := int(*req.Msg.CurrentPage)
	coupons, pageRes, domainErr := ac.couponUseCase.GetCustomCouponList(types.NewPageQuery(&page, &perPage))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var resCoupons []*shared.Coupon
	for _, coupon := range coupons {
		cpn := EntityToResCoupon(coupon)
		resCoupons = append(resCoupons, cpn)
	}

	result := &admin.CouponListResponse{
		Coupons: resCoupons,
		PageResponse: &shared.PageResponse{
			TotalCount:  uint32(pageRes.TotalCount),
			CurrentPage: uint32(pageRes.CurrentPage),
			PerPage:     uint32(pageRes.PerPage),
			TotalPage:   uint32(pageRes.TotalPage),
		},
	}
	return connect.NewResponse(result), nil
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

func EntityToResCoupon(entity *entity.Coupon) *shared.Coupon {
	return &shared.Coupon{
		ID:                entity.ID.String(),
		Name:              entity.Name,
		CouponType:        shared.CouponType(entity.CouponType.ToInt()),
		DiscountAmount:    uint32(entity.DiscountAmount),
		ExpireAt:          timestamppb.New(entity.ExpireAt),
		IsCombinationable: entity.IsCombinationable,
		CreateAt:          timestamppb.New(entity.CreateAt),
	}
}
