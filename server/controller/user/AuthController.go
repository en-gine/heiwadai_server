package user

import (
	"context"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	usecase "server/core/usecase/user"

	"github.com/Songmu/go-httpdate"
	connect_go "github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	authUseCase usecase.AuthUsecase
}

var _ userv1connect.AuthControllerClient = &AuthController{}

func NewAuthController(authUsecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{
		authUseCase: *authUsecase,
	}
}
func (ac *AuthController) Register(ctx context.Context, req *connect_go.Request[user.UserRegisterRequest]) (*connect_go.Response[emptypb.Empty], error) {
	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, err
	}

	_, err = ac.authUseCase.Register(
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
		msg.AcceptTerm,
	)
	return connect_go.NewResponse(&emptypb.Empty{}), err
}

func (ac *AuthController) SignUp(ctx context.Context, req *connect_go.Request[user.UserAuthRequest]) (*connect_go.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.SignUp(msg.Email, msg.Password)
	return connect_go.NewResponse(&emptypb.Empty{}), err
}
func (ac *AuthController) SignIn(ctx context.Context, req *connect_go.Request[user.UserAuthRequest]) (*connect_go.Response[user.UserAuthResponse], error) {
	msg := req.Msg
	token, err := ac.authUseCase.SignIn(msg.Email, msg.Password)
	return connect_go.NewResponse(&user.UserAuthResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), err
}
func (ac *AuthController) ResetPasswordMail(ctx context.Context, req *connect_go.Request[user.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	return connect_go.NewResponse(&emptypb.Empty{}), err
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect_go.Request[user.UpdatePasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.UpdatePassword(msg.Password, msg.Token)
	return connect_go.NewResponse(&emptypb.Empty{}), err
}
func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect_go.Request[user.UpdateEmailRequest]) (*connect_go.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.UpdateEmail(msg.Email, msg.Token)
	return connect_go.NewResponse(&emptypb.Empty{}), err
}
