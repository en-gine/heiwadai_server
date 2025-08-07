package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	"server/core/entity"
	usecase "server/core/usecase/admin"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreController struct {
	storeUseCase usecase.StoreUsecase
}

var _ adminv1connect.StoreControllerClient = &StoreController{}

func NewStoreController(storeUsecase *usecase.StoreUsecase) *StoreController {
	return &StoreController{
		storeUseCase: *storeUsecase,
	}
}

func (ac *StoreController) GetByID(ctx context.Context, req *connect.Request[admin.StoreIDRequest]) (*connect.Response[admin.Store], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	store, domaiErr := ac.storeUseCase.GetStoreByID(storeID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	resStore := StoreToResponse(store)

	return connect.NewResponse(resStore), nil
}

func (ac *StoreController) GetActiveAll(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[admin.Stores], error) {
	stores, domaiErr := ac.storeUseCase.GetActiveList()

	var resStores []*admin.Store
	for _, store := range stores {
		res := StoreToResponse(store)
		resStores = append(resStores, res)
	}
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&admin.Stores{
		Stores: resStores,
	}), nil
}

func (ac *StoreController) GetAll(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[admin.Stores], error) {
	stores, domaiErr := ac.storeUseCase.GetList()

	var resStores []*admin.Store
	for _, store := range stores {
		res := StoreToResponse(store)
		resStores = append(resStores, res)
	}
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&admin.Stores{
		Stores: resStores,
	}), nil
}

func (ac *StoreController) Register(ctx context.Context, req *connect.Request[admin.StoreRegisterRequest]) (*connect.Response[admin.Store], error) {
	store, domaiErr := ac.storeUseCase.Create(
		req.Msg.Name,
		req.Msg.BranchName,
		req.Msg.ZipCode,
		req.Msg.Address,
		req.Msg.Tel,
		req.Msg.SiteURL,
		req.Msg.StampImageData,
		req.Msg.Stayable,
		&req.Msg.StayableInfo.Parking,
		&req.Msg.StayableInfo.Latitude,
		&req.Msg.StayableInfo.Longitude,
		&req.Msg.StayableInfo.AccessInfo,
		&req.Msg.StayableInfo.RestAPIURL,
		&req.Msg.StayableInfo.BookingSystemID,
		&req.Msg.StayableInfo.BookingSystemLoginId,
		&req.Msg.StayableInfo.BookingSystemPassword,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	resStore := StoreToResponse(store)
	return connect.NewResponse(resStore), nil
}

func (ac *StoreController) Update(ctx context.Context, req *connect.Request[admin.StoreUpdateRequest]) (*connect.Response[admin.Store], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	qrcode, err := uuid.Parse(msg.QRCode)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("QRコードが正しい形式ではありません。"))
	}
	unlimitqrcode, err := uuid.Parse(msg.UnLimitedQRCode)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("制限なしQRコードが正しい形式ではありません。"))
	}
	store, domaiErr := ac.storeUseCase.Update(
		storeID,
		req.Msg.Name,
		req.Msg.BranchName,
		req.Msg.ZipCode,
		req.Msg.Address,
		req.Msg.Tel,
		req.Msg.SiteURL,
		req.Msg.StampImageData,
		req.Msg.Stayable,
		&req.Msg.StayableInfo.Parking,
		&req.Msg.StayableInfo.Latitude,
		&req.Msg.StayableInfo.Longitude,
		&req.Msg.StayableInfo.AccessInfo,
		&req.Msg.StayableInfo.RestAPIURL,
		&req.Msg.StayableInfo.BookingSystemID,
		&req.Msg.StayableInfo.BookingSystemLoginId,
		&req.Msg.StayableInfo.BookingSystemPassword,
		req.Msg.IsActive,
		qrcode,
		unlimitqrcode,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	resStore := StoreToResponse(store)
	return connect.NewResponse(resStore), nil
}

func (ac *StoreController) RegenQRCode(ctx context.Context, req *connect.Request[admin.StoreIDRequest]) (*connect.Response[admin.QRResponse], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	qr, domainErr := ac.storeUseCase.RegenQR(storeID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&admin.QRResponse{
		QRCode: qr.String(),
	}), nil
}

func (ac *StoreController) RegenUnlimitQRCode(ctx context.Context, req *connect.Request[admin.StoreIDRequest]) (*connect.Response[admin.UnlimitQRResponse], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}
	qr, domainErr := ac.storeUseCase.RegenUnlimitQR(storeID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&admin.UnlimitQRResponse{
		UnlimitQRCode: qr.String(),
	}), nil
}

func StoreToResponse(store *entity.Store) *admin.Store {
	var resInfo *admin.StayableStoreInfo = nil
	if store.StayableStoreInfo != nil {
		resInfo = StayableInfoToResponse(store.StayableStoreInfo)
	}
	return &admin.Store{
		ID:              store.ID.String(),
		Name:            store.Name,
		BranchName:      store.BranchName,
		ZipCode:         store.ZipCode,
		Address:         store.Address,
		Tel:             store.Tel,
		SiteURL:         store.SiteURL,
		StampImageURL:   store.StampImageURL,
		Stayable:        store.Stayable,
		IsActive:        store.IsActive,
		QRCode:          store.QRCode.String(),
		UnLimitedQRCode: store.UnLimitedQRCode.String(),
		StayableInfo:    resInfo,
	}
}

func StayableInfoToResponse(info *entity.StayableStoreInfo) *admin.StayableStoreInfo {
	return &admin.StayableStoreInfo{
		Parking:               info.Parking,
		Latitude:              info.Latitude,
		Longitude:             info.Longitude,
		AccessInfo:            info.AccessInfo,
		RestAPIURL:            info.RestAPIURL,
		BookingSystemID:       info.BookingSystemID,
		BookingSystemLoginId:  info.BookingSystemLoginID,
		BookingSystemPassword: info.BookingSystemPassword,
	}
}
