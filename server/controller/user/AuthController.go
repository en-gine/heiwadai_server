package user

import (
	"context"
	"errors"
	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	usecase "server/core/usecase/user"
	"server/infrastructure/logger"

	"github.com/Songmu/go-httpdate"
	connect "github.com/bufbuild/connect-go"
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
func (ac *AuthController) Register(ctx context.Context, req *connect.Request[user.UserRegisterRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	birth, err := httpdate.Str2Time(msg.BirthDate, nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("誕生日の形式が不正です"))
	}

	_, domaiErr := ac.authUseCase.Register(
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
	return connect.NewResponse(&emptypb.Empty{}), domaiErr
}

func (ac *AuthController) SignUp(ctx context.Context, req *connect.Request[user.UserAuthRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアップに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
func (ac *AuthController) SignIn(ctx context.Context, req *connect.Request[user.UserAuthRequest]) (*connect.Response[user.UserAuthResponse], error) {
	msg := req.Msg
	token, err := ac.authUseCase.SignIn(msg.Email, msg.Password)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&user.UserAuthResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}
func (ac *AuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[user.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect.Request[user.UpdatePasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.UpdatePassword(msg.Password, msg.Token)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect.Request[user.UpdateEmailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.UpdateEmail(msg.Email, msg.Token)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
