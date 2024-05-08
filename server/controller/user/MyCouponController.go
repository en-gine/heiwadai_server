package user

import (
	"context"
	"errors"

	"server/api/v1/shared"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/core/entity"
	usecase "server/core/usecase/user"
	"server/router"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MyCouponController struct {
	couponUseCase usecase.UserAttachedCouponUsecase
}

var _ userv1connect.MyCouponControllerClient = &MyCouponController{}

func NewMyCouponController(couponUsecase *usecase.UserAttachedCouponUsecase) *MyCouponController {
	return &MyCouponController{
		couponUseCase: *couponUsecase,
	}
}

func (ac *MyCouponController) GetDetail(ctx context.Context, req *connect.Request[user.CouponIDRequest]) (*connect.Response[shared.Coupon], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	couponID, err := uuid.Parse(req.Msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	coupon, domaiErr := ac.couponUseCase.GetByID(userID, couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var TargetStores []*shared.Store
	for _, store := range coupon.Coupon.TargetStore {
		TargetStores = append(TargetStores, &shared.Store{
			ID:   store.ID.String(),
			Name: store.Name,
		})
	}
	result := &shared.Coupon{
		ID:                coupon.ID.String(),
		Name:              coupon.Coupon.Name,
		CouponType:        shared.CouponType(coupon.Coupon.CouponType.ToInt()),
		DiscountAmount:    uint32(coupon.Coupon.DiscountAmount),
		ExpireAt:          timestamppb.New(coupon.Coupon.ExpireAt),
		IsCombinationable: coupon.Coupon.IsCombinationable,
		TargetStore:       TargetStores,
		Notices:           coupon.Coupon.Notices,
		CreateAt:          timestamppb.New(coupon.Coupon.CreateAt),
	}

	return connect.NewResponse(result), nil
}

func (ac *MyCouponController) GetList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.MyCouponsResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	entities, domaiErr := ac.couponUseCase.GetMyList(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	resposnse := UserAttachedCouponToResponse(entities)

	return connect.NewResponse(resposnse), nil
}

func (ac *MyCouponController) Use(ctx context.Context, req *connect.Request[user.CouponIDRequest]) (*connect.Response[emptypb.Empty], error) {
	id := req.Msg.ID
	couponID, err := uuid.Parse(id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}
	domaiErr := ac.couponUseCase.UseMyCoupon(userID, couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *MyCouponController) GetExpiredList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.MyCouponsResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	entities, domaiErr := ac.couponUseCase.GetMyExpireList(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	resposnse := UserAttachedCouponToResponse(entities)

	return connect.NewResponse(resposnse), nil
}

func (ac *MyCouponController) GetUsedList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.MyCouponsResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}

	entities, domaiErr := ac.couponUseCase.GetMyUsedList(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	resposnse := UserAttachedCouponToResponse(entities)

	return connect.NewResponse(resposnse), nil

}

func UserAttachedCouponToResponse(entities []*entity.UserAttachedCoupon) *user.MyCouponsResponse {
	var response []*shared.Coupon
	for _, coupon := range entities {
		response = append(response, &shared.Coupon{
			ID:                coupon.ID.String(),
			Name:              coupon.Coupon.Name,
			CouponType:        shared.CouponType(coupon.Coupon.CouponType.ToInt()),
			DiscountAmount:    uint32(coupon.Coupon.DiscountAmount),
			ExpireAt:          timestamppb.New(coupon.Coupon.ExpireAt),
			IsCombinationable: coupon.Coupon.IsCombinationable,
			CreateAt:          timestamppb.New(coupon.Coupon.CreateAt),
		})
	}

	return &user.MyCouponsResponse{
		Coupons: response,
	}
}
