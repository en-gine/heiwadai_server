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

func (ac *AdminCouponController) GetUserCouponList(ctx context.Context, req *connect.Request[admin.UserCouponRequest]) (*connect.Response[admin.UserAttachedCouponsResponse], error) {
	userID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	perPage := int(*req.Msg.Pager.PerPage)
	page := int(*req.Msg.Pager.CurrentPage)

	pager := types.NewPageQuery(&perPage, &page)

	entities, pageRes, domaiErr := ac.couponUseCase.GetUsersCouponList(userID, pager)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	var resPage *shared.PageResponse

	if pageRes != nil {
		resPage = &shared.PageResponse{
			TotalCount:  uint32(pageRes.TotalCount),
			CurrentPage: uint32(pageRes.CurrentPage),
			PerPage:     uint32(pageRes.PerPage),
			TotalPage:   uint32(pageRes.TotalPage),
		}
	} else {
		resPage = &shared.PageResponse{
			TotalCount:  0,
			CurrentPage: 0,
			PerPage:     0,
			TotalPage:   0,
		}
	}
	var attachedCoupons []*shared.UserAttachedCoupon
	for _, entity := range entities {
		cpn := EntityToCoupon(entity.Coupon)

		atcCpn := &shared.UserAttachedCoupon{
			UserID: userID.String(),
			Coupon: cpn,
			UsedAt: timestamppb.New(entity.ExpireAt),
		}
		attachedCoupons = append(attachedCoupons, atcCpn)
	}
	result := &admin.UserAttachedCouponsResponse{
		UserAttachedCoupons: attachedCoupons,
		PageResponse:        resPage,
	}

	return connect.NewResponse(result), nil
}

func (ac *AdminCouponController) CreateCustomCoupon(ctx context.Context, req *connect.Request[admin.CreateCustomCouponRequest]) (*connect.Response[shared.CustomCoupon], error) {
	var TargetStores []*entity.Store
	for _, id := range req.Msg.TargetStoresID {
		storeId, err := uuid.Parse(id)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("StoreのUUIDが正しい形式ではありません。"))
		}

		TargetStores = append(TargetStores, &entity.Store{
			ID: storeId,
		})
	}

	entity, domaiErr := ac.couponUseCase.CreateCustomCoupon(
		req.Msg.Name,
		uint(req.Msg.DiscountAmount),
		req.Msg.ExpireAt.AsTime(),
		req.Msg.IsCombinationable,
		req.Msg.Notices,
		TargetStores,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	cpn := EntityToCustomCoupon(entity)

	return connect.NewResponse(cpn), nil
}

func (ac *AdminCouponController) GetCustomCouponByID(ctx context.Context, req *connect.Request[admin.CouponIDRequest]) (*connect.Response[shared.CustomCoupon], error) {
	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	entity, domaiErr := ac.couponUseCase.GetCustomCouponByID(couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	cpn := EntityToCustomCoupon(entity)

	return connect.NewResponse(cpn), nil
}

func (ac *AdminCouponController) GetDefaultEmptyCoupon(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[admin.DefaultCouponResponse], error) {
	coupon, err := ac.couponUseCase.DefaultEmptyCustomCoupon()
	if err != nil {
		return nil, controller.ErrorHandler(err)
	}

	cpn := EntityToCustomCoupon(coupon)

	return connect.NewResponse(
		&admin.DefaultCouponResponse{
			Coupon: cpn,
		},
	), nil
}

func (ac *AdminCouponController) GetCustomCouponList(ctx context.Context, req *connect.Request[shared.Pager]) (*connect.Response[admin.CouponListResponse], error) {
	perPage := int(*req.Msg.PerPage)
	page := int(*req.Msg.CurrentPage)
	coupons, pageRes, domainErr := ac.couponUseCase.GetCustomCouponList(types.NewPageQuery(&page, &perPage))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var resCoupons []*shared.CustomCoupon
	for _, coupon := range coupons {
		cpn := EntityToCustomCoupon(coupon)
		resCoupons = append(resCoupons, cpn)
	}
	var resPage *shared.PageResponse
	if pageRes != nil {
		resPage = &shared.PageResponse{
			TotalCount:  uint32(pageRes.TotalCount),
			CurrentPage: uint32(pageRes.CurrentPage),
			PerPage:     uint32(pageRes.PerPage),
			TotalPage:   uint32(pageRes.TotalPage),
		}
	} else {
		resPage = &shared.PageResponse{
			TotalCount:  0,
			CurrentPage: 0,
			PerPage:     0,
			TotalPage:   0,
		}
	}
	result := &admin.CouponListResponse{
		Coupons:      resCoupons,
		PageResponse: resPage,
	}
	return connect.NewResponse(result), nil
}

func (ac *AdminCouponController) SaveCustomCoupon(ctx context.Context, req *connect.Request[admin.SaveCustomCouponRequest]) (*connect.Response[emptypb.Empty], error) {
	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	var TargetStores []*entity.Store
	for _, id := range req.Msg.TargetStoresID {
		storeId, err := uuid.Parse(id)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("StoreのUUIDが正しい形式ではありません。"))
		}

		TargetStores = append(TargetStores, &entity.Store{
			ID: storeId,
		})
	}
	domaiErr := ac.couponUseCase.SaveCustomCoupon(couponID, req.Msg.Name, uint(req.Msg.DiscountAmount), req.Msg.ExpireAt.AsTime(), req.Msg.IsCombinationable, req.Msg.Notices, TargetStores)
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

func EntityToCustomCoupon(entity *entity.Coupon) *shared.CustomCoupon {
	var stores []*shared.Store
	for _, store := range entity.TargetStore {
		stores = append(stores, &shared.Store{
			ID: store.ID.String(),
		})
	}
	return &shared.CustomCoupon{
		ID:                entity.ID.String(),
		Name:              entity.Name,
		CouponType:        shared.CouponType(entity.CouponType.ToInt()),
		TargetStore:       stores,
		DiscountAmount:    uint32(entity.DiscountAmount),
		ExpireAt:          timestamppb.New(entity.ExpireAt),
		Status:            shared.CouponStatus(entity.Status),
		IsCombinationable: entity.IsCombinationable,
		Notices:           entity.Notices,
		CreateAt:          timestamppb.New(entity.CreateAt),
	}
}

func EntityToCoupon(entity *entity.Coupon) *shared.Coupon {
	var stores []*shared.Store
	for _, store := range entity.TargetStore {
		stores = append(stores, &shared.Store{
			ID: store.ID.String(),
		})
	}
	return &shared.Coupon{
		ID:                entity.ID.String(),
		Name:              entity.Name,
		CouponType:        shared.CouponType(entity.CouponType.ToInt()),
		DiscountAmount:    uint32(entity.DiscountAmount),
		TargetStore:       stores,
		ExpireAt:          timestamppb.New(entity.ExpireAt),
		IsCombinationable: entity.IsCombinationable,
		Notices:           entity.Notices,
		CreateAt:          timestamppb.New(entity.CreateAt),
	}
}
