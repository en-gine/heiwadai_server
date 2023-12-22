package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	usecase "server/core/usecase/admin"
	"server/infrastructure/logger"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AnonAuthController struct {
	authUseCase usecase.AuthUsecase
}

var _ adminv1connect.AnonAuthControllerClient = &AnonAuthController{}

func NewAnonAuthController(authUsecase *usecase.AuthUsecase) *AnonAuthController {
	return &AnonAuthController{
		authUseCase: *authUsecase,
	}
}

func (ac *AnonAuthController) Register(ctx context.Context, req *connect.Request[admin.AdminRegisterRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	_, domaiErr := ac.authUseCase.Register(
		msg.Name,
		uuid.MustParse(msg.BelongStoreID),
		msg.Mail,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignUp(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアップに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignIn(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[admin.AnonAuthTokenResponse], error) {
	msg := req.Msg
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password)
	if domainErr != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインインに失敗しました。"))
	}
	return connect.NewResponse(&admin.AnonAuthTokenResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}

func (ac *AnonAuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[admin.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("リセットメールの送信に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
