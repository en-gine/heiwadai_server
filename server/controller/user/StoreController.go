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

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreController struct {
	storeUseCase usecase.StoreUsecase
}

var _ userv1connect.StoreControllerClient = &StoreController{}

func NewStoreController(storeUsecase *usecase.StoreUsecase) *StoreController {
	return &StoreController{
		storeUseCase: *storeUsecase,
	}
}

func (ac *StoreController) GetByID(ctx context.Context, req *connect.Request[user.StoreIDRequest]) (*connect.Response[shared.Store], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("UUIDが正しい形式ではありません。"))
	}

	store, domaiErr := ac.storeUseCase.GetByID(storeID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	resposnse := StoreToResponse(store)

	return connect.NewResponse(resposnse), nil
}

func (ac *StoreController) GetAll(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[shared.Stores], error) {
	stores, domaiErr := ac.storeUseCase.GetAll()

	var resStores []*shared.Store
	for _, store := range stores {
		res := StoreToResponse(store)
		resStores = append(resStores, res)
	}
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&shared.Stores{
		Stores: resStores,
	}), nil
}

func (ac *StoreController) GetStayableByID(ctx context.Context, req *connect.Request[user.StoreIDRequest]) (*connect.Response[shared.StayableStore], error) {
	msg := req.Msg
	storeID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, errors.New("UUIDが正しい形式ではありません。")
	}

	stayableStore, domaiErr := ac.storeUseCase.GetStayableByID(storeID)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	store := StoreToResponse(stayableStore.Store)
	info := StayableInfoToResponse(stayableStore.StayableStoreInfo)

	return connect.NewResponse(&shared.StayableStore{
		Store: store,
		Info:  info,
	}), nil
}

func (ac *StoreController) GetStayables(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[shared.StayableStores], error) {
	stayables, domaiErr := ac.storeUseCase.GetStayables()

	var resStores []*shared.StayableStore
	for _, stays := range stayables {
		store := StoreToResponse(stays.Store)
		info := StayableInfoToResponse(stays.StayableStoreInfo)

		resStores = append(resStores, &shared.StayableStore{
			Store: store,
			Info:  info,
		})
	}
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&shared.StayableStores{
		StayableStores: resStores,
	}), nil
}

func StoreToResponse(store *entity.Store) *shared.Store {
	return &shared.Store{
		ID:                store.ID.String(),
		Name:              store.Name,
		BranchName:        store.BranchName,
		ZipCode:           store.ZipCode,
		Address:           store.Address,
		Tel:               store.Tel,
		SiteURL:           store.SiteURL,
		StampImageURL:     store.StampImageURL,
		Stayable:          store.Stayable,
		IsActive:          store.IsActive,
		QRCode:            store.QRCode.String(),
		UnLimitedQRCode:   store.UnLimitedQRCode.String(),
		StayableStoreInfo: StayableInfoToResponse(store.StayableStoreInfo),
	}
}

func StayableInfoToResponse(info *entity.StayableStoreInfo) *shared.StayableStoreInfo {
	if info == nil {
		return &shared.StayableStoreInfo{}
	}
	return &shared.StayableStoreInfo{
		Parking:         info.Parking,
		Latitude:        info.Latitude,
		Longitude:       info.Longitude,
		AccessInfo:      info.AccessInfo,
		RestAPIURL:      info.RestAPIURL,
		BookingSystemID: info.BookingSystemID,
	}
}
