package admin

import (
	"context"

	"server/api/v1/admin"
	adminv1 "server/api/v1/admin"
	"server/api/v1/shared"
	userv1 "server/api/v1/user"

	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	userController "server/controller/user"

	"server/controller/util"
	"server/core/entity"
	"server/core/infra/queryService/types"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserDataController struct {
	userUsecase    usecase.UserDataUsecase
	checkinUsecase usecase.UserCheckinUsecase
}

var _ adminv1connect.UserDataControllerClient = &UserDataController{}

func NewUserDataController(userUsecase *usecase.UserDataUsecase, checkinUsecase *usecase.UserCheckinUsecase) *UserDataController {
	return &UserDataController{
		userUsecase:    *userUsecase,
		checkinUsecase: *checkinUsecase,
	}
}

func (u *UserDataController) Update(ctx context.Context, req *connect.Request[adminv1.UserUpdateDataRequest]) (*connect.Response[adminv1.UserDataResponse], error) {
	msg := req.Msg
	user := msg.User

	entity, domainErr := u.userUsecase.Update(
		uuid.MustParse(user.ID),
		user.FirstName,
		user.LastName,
		user.FirstNameKana,
		user.LastNameKana,
		user.CompanyName,
		util.TimeStampPtrToTimePtr(user.BirthDate),
		user.ZipCode,
		int(user.Prefecture),
		user.City,
		user.Address,
		user.Tel,
		user.Mail,
		user.AcceptMail,
		msg.InnerNote,
		msg.IsBlackCustomer,
	)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	res := connect.NewResponse(&adminv1.UserDataResponse{
		User: &userv1.UserDataResponse{
			ID:            entity.User.ID.String(),
			FirstName:     entity.User.FirstName,
			LastName:      entity.User.LastName,
			FirstNameKana: entity.User.FirstNameKana,
			LastNameKana:  entity.User.LastNameKana,
			CompanyName:   entity.User.CompanyName,
			BirthDate:     util.TimePtrToTimeStampPtr(entity.User.BirthDate),
			ZipCode:       entity.User.ZipCode,
			Prefecture:    shared.Prefecture(entity.User.Prefecture.ToInt()),
			City:          entity.User.City,
			Address:       entity.User.Address,
			Tel:           entity.User.Tel,
			AcceptMail:    entity.User.AcceptMail,
		},
		InnerNote:       entity.UserOption.InnerNote,
		IsBlackCustomer: entity.UserOption.IsBlackCustomer,
	})
	return res, nil
}

func (u *UserDataController) Delete(ctx context.Context, req *connect.Request[adminv1.UserDeleteRequest]) (*connect.Response[emptypb.Empty], error) {
	domainErr := u.userUsecase.Delete(uuid.MustParse(req.Msg.ID))
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (u *UserDataController) GetList(ctx context.Context, req *connect.Request[adminv1.UserListFilterRequest]) (*connect.Response[admin.UserListResponse], error) {
	msg := req.Msg
	filter := msg.Search

	var pref *entity.Prefecture = nil
	if pref != nil {
		tmp := entity.Prefecture(*filter.Prefecture)
		pref = &tmp
	}

	query := &types.UserQuery{
		FirstName:     filter.FirstName,
		LastName:      filter.LastName,
		FirstNameKana: filter.FirstNameKana,
		LastNameKana:  filter.LastNameKana,
		Prefecture:    pref,
	}

	pager := &types.PageQuery{}

	users, pageResponse, domainErr := u.userUsecase.GetList(query, pager)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	var res *admin.UserListResponse
	var resUsers []*admin.UserWithCheckIn
	var resPage *shared.PageResponse
	for _, user := range users {
		resUsers = append(resUsers,
			&admin.UserWithCheckIn{
				User:          userController.UserEntityToResponse(user.User),
				LastCheckinAt: util.TimePtrToTimeStampPtr(user.LastCheckinAt),
			},
		)
	}

	resPage = &shared.PageResponse{
		TotalCount:  uint32(pageResponse.TotalCount),
		CurrentPage: uint32(pageResponse.CurrentPage),
		PerPage:     uint32(pageResponse.PerPage),
		TotalPage:   uint32(pageResponse.TotalPage),
	}
	res = &adminv1.UserListResponse{
		Users:        resUsers,
		PageResponse: resPage,
	}
	return connect.NewResponse(res), nil
}
