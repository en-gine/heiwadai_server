package user

import (
	"context"
	"errors"
	"fmt"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/controller/util"
	usecase "server/core/usecase/user"
	"server/infrastructure/env"

	connect "github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AnonAuthController struct {
	authUseCase usecase.AuthUsecase
}

var _ userv1connect.AnonAuthControllerClient = &AnonAuthController{}

func NewAnonAuthController(authUsecase *usecase.AuthUsecase) *AnonAuthController {
	return &AnonAuthController{
		authUseCase: *authUsecase,
	}
}

func (ac *AnonAuthController) Register(ctx context.Context, req *connect.Request[user.UserRegisterRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	_, domaiErr := ac.authUseCase.Register(
		msg.FirstName,
		msg.LastName,
		msg.FirstNameKana,
		msg.LastNameKana,
		msg.CompanyName,
		util.TimeStampPtrToTimePtr(msg.BirthDate),
		msg.ZipCode,
		int(msg.Prefecture),
		msg.City,
		msg.Address,
		msg.Tel,
		msg.Mail,
		msg.AcceptMail,
		msg.AcceptTerm,
	)
	if domaiErr != nil {
		return nil, controller.ErrorHandler(domaiErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignUp(ctx context.Context, req *connect.Request[user.UserAuthRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	_, domainErr := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignIn(ctx context.Context, req *connect.Request[user.UserAuthRequest]) (*connect.Response[user.AnonTokenResponse], error) {
	remoteAddr := req.Peer().Addr
	userAgent := req.Header().Get("User-Agent")
	if env.GetEnv(env.EnvMode) == "dev" {
		fmt.Println(ctx, req)
	}
	msg := req.Msg
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password, remoteAddr, userAgent)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&user.AnonTokenResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}

func (ac *AnonAuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[user.UserMailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	domainErr := ac.authUseCase.ResetPasswordMail(msg.Email)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) IsUnderRegister(ctx context.Context, req *connect.Request[user.UserMailRequest]) (*connect.Response[user.IsUnderRegisterResponse], error) {
	msg := req.Msg
	isYes, domainErr := ac.authUseCase.IsUnderRegister(msg.Email)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&user.IsUnderRegisterResponse{
		IsUnderRegister: isYes,
	}), nil
}

func (ac *AnonAuthController) SetNewPassword(ctx context.Context, req *connect.Request[user.SetNewPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	err := ac.authUseCase.UpdatePassword(msg.Password, msg.AccessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("パスワードの変更に失敗しました。\nネットワークの問題や同じパスワードに変更した、などの理由が考えられます。"))
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) ResendInviteMail(ctx context.Context, req *connect.Request[user.UserMailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	domainErr := ac.authUseCase.ResendInviteMail(msg.Email)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) DeleteUnderRegisterUser(ctx context.Context, req *connect.Request[user.UserMailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	domainErr := ac.authUseCase.DeleteUnderRegisterUser(msg.Email)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
