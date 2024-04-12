package user

import (
	"context"
	"errors"

	"server/api/v1/shared"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"
	"server/router"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CheckInController struct {
	checkInUseCase usecase.UserCheckinUsecase
}

var _ userv1connect.CheckinControllerClient = &CheckInController{}

func NewCheckInController(checkinUsecase *usecase.UserCheckinUsecase) *CheckInController {
	return &CheckInController{
		checkInUseCase: *checkinUsecase,
	}
}

func (ac *CheckInController) GetStampCard(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[user.StampCardResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}
	entity, domaiErr := ac.checkInUseCase.GetStampCard(userID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	stamps := []*user.CheckinStamp{}

	for _, entity := range entity.Stamps {
		var stamp *user.CheckinStamp

		if entity.CheckinID != nil {
			strChkID := entity.CheckinID.String()
			strStoreID := entity.StoreID.String()
			stamp = &user.CheckinStamp{
				ID:              &strChkID,
				StoreName:       entity.StoreName,
				StoreID:         &strStoreID,
				StoreStampImage: entity.StoreStampImage,
				CheckInAt:       timestamppb.New(*entity.CheckInAt),
			}
		}
		stamps = append(stamps, stamp)
	}

	response := &user.StampCardResponse{
		Stamps: stamps,
	}

	return connect.NewResponse(response), nil
}

func (ac *CheckInController) Checkin(ctx context.Context, req *connect.Request[user.CheckinRequest]) (*connect.Response[user.CheckinResponse], error) {
	if ctx.Value(router.UserIDKey) == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。"))
	}

	userID, err := uuid.Parse(ctx.Value(router.UserIDKey).(string))
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("ユーザーIDが取得できませんでした。UUIDの形式が不正です。"))
	}
	qrHash, err := uuid.Parse(req.Msg.QrHash)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("QRコードの値が正しくありません。"))
	}

	mayCoupon, domaiErr := ac.checkInUseCase.Checkin(userID, qrHash)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var TargetStores []*shared.Store
	for _, store := range mayCoupon.Coupon.TargetStore {
		TargetStores = append(TargetStores, &shared.Store{
			ID:   store.ID.String(),
			Name: store.Name,
		})
	}

	coupon := &shared.Coupon{
		ID:                mayCoupon.ID.String(),
		Name:              mayCoupon.Coupon.Name,
		CouponType:        shared.CouponType(mayCoupon.Coupon.CouponType.ToInt()),
		DiscountAmount:    uint32(mayCoupon.Coupon.DiscountAmount),
		ExpireAt:          timestamppb.New(mayCoupon.Coupon.ExpireAt),
		IsCombinationable: mayCoupon.Coupon.IsCombinationable,
		TargetStore:       TargetStores,
		Notices:           mayCoupon.Coupon.Notices,
	}

	response := &user.CheckinResponse{
		MayCoupon: coupon,
	}

	return connect.NewResponse(response), nil
}
