package user

import (
	userv1 "server/api/v1/user"
	"server/core/infra/repository"
	usecase "server/core/usecase/user"

	"github.com/Songmu/go-httpdate"
	connect "github.com/bufbuild/connect-go"
)

type UserDataController struct {
	userRepository repository.IUserRepository
}

func NewUserDataController(userRepository repository.IUserRepository) *UserDataController {
	return &UserDataController{
		userRepository: userRepository,
	}
}

func (u *UserDataController) Update(req *connect.Request[userv1.UserDataRequest]) (*connect.Response[userv1.UserDataResponse], error) {
	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, err
	}

	service := usecase.NewUserDataUsecase(u.userRepository)
	user, err := service.Update(
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
	res := connect.NewResponse(&userv1.UserDataResponse{
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
