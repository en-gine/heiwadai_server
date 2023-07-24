package user

import (
	"context"
	"errors"
	userv1 "server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	usecase "server/core/usecase/user"

	"github.com/Songmu/go-httpdate"
	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

type UserDataController struct {
	usecase usecase.UserDataUsecase
}

var _ userv1connect.UserDataControllerClient = &UserDataController{}

func NewUserDataController(userusecase *usecase.UserDataUsecase) *UserDataController {
	return &UserDataController{
		usecase: *userusecase,
	}
}

func (u *UserDataController) Update(ctx context.Context, req *connect.Request[userv1.UserUpdateDataRequest]) (*connect.Response[userv1.UserDataResponse], error) {
	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("誕生日の形式が不正です"))
	}

	user, domainErr := u.usecase.Update(
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
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
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
