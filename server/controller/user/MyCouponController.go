package user

import (
	"context"
	"errors"
	"server/api/v1/shared"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MyCouponController struct {
	authUseCase   usecase.AuthUsecase
	couponUseCase usecase.UserAttachedCouponUsecase
}

var _ userv1connect.MyCouponControllerClient = &MyCouponController{}

func MyNewCouponController(authUsecase *usecase.AuthUsecase, couponUsecase *usecase.UserAttachedCouponUsecase) *MyCouponController {
	return &MyCouponController{
		authUseCase:   *authUsecase,
		couponUseCase: *couponUsecase,
	}
}
func (ac *MyCouponController) GetDetail(ctx context.Context, req *connect.Request[user.CouponIDRequest]) (*connect.Response[user.Coupon], error) {
	userID := ctx.Value("userID").(uuid.UUID)

	if userID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
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
	result := &user.Coupon{
		ID:                coupon.ID.String(),
		Name:              coupon.Coupon.Name,
		CouponType:        user.CouponType(coupon.Coupon.CouponType.ToInt()),
		DiscountAmount:    uint32(coupon.Coupon.DiscountAmount),
		ExpireAt:          timestamppb.New(coupon.Coupon.ExpireAt),
		IsCombinationable: coupon.Coupon.IsCombinationable,
		TargetStore:       TargetStores,
		Notices:           coupon.Coupon.Notices,
	}

	return connect.NewResponse(result), nil
}

func (ac *MyCouponController) GetList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.MyCouponsResponse], error) {
	userID := ctx.Value("userID").(uuid.UUID)
	if userID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	entities, domaiErr := ac.couponUseCase.GetMyList(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var response []*user.Coupon
	for _, coupon := range entities {
		response = append(response, &user.Coupon{
			ID:                coupon.ID.String(),
			Name:              coupon.Coupon.Name,
			CouponType:        user.CouponType(coupon.Coupon.CouponType.ToInt()),
			DiscountAmount:    uint32(coupon.Coupon.DiscountAmount),
			ExpireAt:          timestamppb.New(coupon.Coupon.ExpireAt),
			IsCombinationable: coupon.Coupon.IsCombinationable,
		})
	}

	resposnse := &user.MyCouponsResponse{
		Coupons: response,
	}

	return connect.NewResponse(resposnse), domaiErr
}

func (ac *MyCouponController) Use(ctx context.Context, req *connect.Request[user.CouponIDRequest]) (*connect.Response[emptypb.Empty], error) {

	id := req.Msg.ID
	couponID, err := uuid.Parse(id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	userID := ctx.Value("userID").(uuid.UUID)
	if userID == uuid.Nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}
	domaiErr := ac.couponUseCase.UseMyCoupon(userID, couponID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
