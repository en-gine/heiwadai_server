package user

import (
	"context"
	userv1 "server/api/v1/user"
	usecase "server/usecase/user"

	"github.com/Songmu/go-httpdate"
	connect "github.com/bufbuild/connect-go"
)

type UserFetcherController struct {
}

func (u *UserFetcherController) Call(ctx context.Context, req *connect.Request[userv1.UserRegisterRequest]) (*connect.Response[userv1.UserRegisterResponse], error) {

	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, err
	}

	service := usecase.RegisterUserUsecase{}
	user, err := service.Exec(
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
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&userv1.UserRegisterResponse{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FirstNameKana: user.FirstNameKana,
		LastNameKana:  user.LastNameKana,
		CompanyName:   *user.CompanyName,
		BirthDate:     user.BirthDate.String(),
		ZipCode:       *user.ZipCode,
		Prefecture:    user.Prefecture,
		City:          *user.City,
		Address:       *user.Address,
		Tel:           *user.Tel,
		Mail:          user.Mail,
		AcceptMail:    user.AcceptMail,
	})
	return res, nil

}
