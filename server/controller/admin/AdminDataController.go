package admin

import (
	"context"

	adminv1 "server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
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

	admin, domainErr := u.usecase.Update(
		uuid.MustParse(msg.ID),
		msg.Name,
		msg.IsActive,
		msg.Mail,
		uuid.MustParse(msg.StoreID),
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	res := connect.NewResponse(&adminv1.AdminDataResponse{
		ID:       admin.ID.String(),
		Name:     admin.Name,
		IsActive: admin.IsActive,
		Mail:     admin.Mail,
		StoreID:  admin.BelongStore.ID.String(),
	})
	return res, nil
}
