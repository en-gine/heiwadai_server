package admin

import (
	"context"
	"errors"
	"fmt"

	adminv1 "server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	"server/core/entity"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminDataController struct {
	usecase usecase.AdminDataUsecase
}

var _ adminv1connect.AdminDataControllerClient = &AdminDataController{}

func NewAdminDataController(adminusecase *usecase.AdminDataUsecase) *AdminDataController {
	return &AdminDataController{
		usecase: *adminusecase,
	}
}

func (u *AdminDataController) Update(ctx context.Context, req *connect.Request[adminv1.AdminUpdateDataRequest]) (*connect.Response[adminv1.AdminDataResponse], error) {
	msg := req.Msg
	fmt.Println(msg.ID)
	adminId, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Idの形式が正しくありません"))
	}
	storeId, err := uuid.Parse(msg.StoreID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("storeIdの形式が正しくありません"))
	}

	admin, domainErr := u.usecase.Update(
		adminId,
		msg.Name,
		msg.IsActive,
		storeId,
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	resAdmin := AdminEntityToResponse(admin)

	res := connect.NewResponse(resAdmin)
	return res, nil
}

func (u *AdminDataController) GetByID(ctx context.Context, req *connect.Request[adminv1.AdminDataRequest]) (*connect.Response[adminv1.AdminDataResponse], error) {
	msg := req.Msg

	adminId, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	entity, domainErr := u.usecase.GetByID(adminId)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	resAdmin := AdminEntityToResponse(entity)

	res := connect.NewResponse(resAdmin)
	return res, nil
}

func (u *AdminDataController) GetAll(ctx context.Context, req *connect.Request[adminv1.AdminListRequest]) (*connect.Response[adminv1.AdminListResponse], error) {
	admins, domainErr := u.usecase.GetAll()
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var resAdmins []*adminv1.AdminDataResponse
	for _, admin := range admins {
		resAdmins = append(resAdmins, AdminEntityToResponse(admin))
	}

	return connect.NewResponse(&adminv1.AdminListResponse{
		Admins: resAdmins,
	}), nil
}

func (u *AdminDataController) Delete(ctx context.Context, req *connect.Request[adminv1.AdminDataRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	adminId, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	domainErr := u.usecase.Delete(adminId)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func AdminEntityToResponse(entity *entity.Admin) *adminv1.AdminDataResponse {
	return &adminv1.AdminDataResponse{
		ID:          entity.ID.String(),
		Name:        entity.Name,
		IsActive:    entity.IsActive,
		Mail:        entity.Mail,
		StoreID:     entity.BelongStore.ID.String(),
		IsConfirmed: entity.IsConfirmed,
	}
}
