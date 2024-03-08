package admin

import (
	"context"
	"errors"

	adminv1 "server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/api/v1/shared"
	"server/controller"
	"server/core/entity"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	adminID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("Idの形式が正しくありません"))
	}
	storeID, err := uuid.Parse(msg.StoreID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("storeIDの形式が正しくありません"))
	}

	admin, domainErr := u.usecase.Update(
		adminID,
		msg.Name,
		msg.IsActive,
		storeID,
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

	adminID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	entity, domainErr := u.usecase.GetByID(adminID)
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

	adminID, err := uuid.Parse(msg.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	domainErr := u.usecase.Delete(adminID)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (u *AdminDataController) GetLoginLogList(ctx context.Context, req *connect.Request[adminv1.AdminLoginLogRequest]) (*connect.Response[adminv1.AdminLoginLogListResponse], error) {
	msg := req.Msg

	userID, err := uuid.Parse(msg.UserID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	perPage := int(*req.Msg.Pager.PerPage)
	curPage := int(*req.Msg.Pager.CurrentPage)

	pager := types.NewPageQuery(&curPage, &perPage)
	logs, pageResponse, domaiErr := u.usecase.GetUserLoginLogList(userID, pager)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	var resLogs []*adminv1.AdminLoginLog
	for _, log := range logs {
		resLogs = append(resLogs, &adminv1.AdminLoginLog{
			UserID:    log.UserID.String(),
			LoginAt:   timestamppb.New(log.LoginAt),
			IP:        log.RemoteID,
			UserAgent: log.UserAgent,
		})
	}
	var resPage *shared.PageResponse

	if pageResponse != nil {
		resPage = &shared.PageResponse{
			TotalCount:  uint32(pageResponse.TotalCount),
			CurrentPage: uint32(pageResponse.CurrentPage),
			PerPage:     uint32(pageResponse.PerPage),
			TotalPage:   uint32(pageResponse.TotalPage),
		}
	} else {
		resPage = &shared.PageResponse{
			TotalCount:  0,
			CurrentPage: 0,
			PerPage:     0,
			TotalPage:   0,
		}
	}
	return connect.NewResponse(&adminv1.AdminLoginLogListResponse{
		LoginLogs:    resLogs,
		PageResponse: resPage,
	}), nil
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
