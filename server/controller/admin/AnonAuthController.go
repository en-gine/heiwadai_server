package admin

import (
	"context"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	"server/controller"
	usecase "server/core/usecase/admin"

	connect "github.com/bufbuild/connect-go"
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

func (ac *AnonAuthController) SignUp(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	_, domainErr := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignIn(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[admin.AnonAuthTokenResponse], error) {
	msg := req.Msg
	remoteAddr := req.Peer().Addr
	userAgent := req.Header().Get("User-Agent")
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password, remoteAddr, userAgent)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&admin.AnonAuthTokenResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}

func (ac *AnonAuthController) SetNewPassword(ctx context.Context, req *connect.Request[admin.SetNewPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	domainErr := ac.authUseCase.UpdatePassword(msg.Password, msg.AccessToken)

	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[admin.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	domainErr := ac.authUseCase.ResetPasswordMail(msg.Email)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
