package user

import (
	"context"
	userv1 "server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	usecase "server/core/usecase/user"

	"github.com/Songmu/go-httpdate"
	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

type UserDataController struct {
	authAction     action.IAuthAction
	userRepository repository.IUserRepository
	userQuery      queryservice.IUserQueryService
}

var _ userv1connect.UserDataControllerClient = &UserDataController{}

func NewUserDataController(authAction action.IAuthAction, userRepository repository.IUserRepository, userQuery queryservice.IUserQueryService) *UserDataController {
	return &UserDataController{
		authAction:     authAction,
		userRepository: userRepository,
		userQuery:      userQuery,
	}
}

func (u *UserDataController) Update(ctx context.Context, req *connect.Request[userv1.UserUpdateDataRequest]) (*connect.Response[userv1.UserDataResponse], error) {
	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, err
	}

	service := usecase.NewUserDataUsecase(u.userRepository, u.userQuery)
	user, err := service.Update(
		uuid.MustParse(msg.ID),
		msg.FirstName,
		msg.LastName,
		msg.FirstNameKana,
		msg.LastNameKana,
		&msg.CompanyName,
		birth,
		&msg.ZipCode,
		msg.Prefecture,
		&msg.City,
		&msg.Address,
		&msg.Tel,
		msg.Mail,
		msg.AcceptMail,
	)

	res := connect.NewResponse(&userv1.UserDataResponse{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FirstNameKana: user.FirstNameKana,
		LastNameKana:  user.LastNameKana,
		CompanyName:   *user.CompanyName,
		BirthDate:     user.BirthDate.String(),
		ZipCode:       *user.ZipCode,
		Prefecture:    user.Prefecture.String(),
		City:          *user.City,
		Address:       *user.Address,
		Tel:           *user.Tel,
		AcceptMail:    user.AcceptMail,
	})
	return res, nil

}
