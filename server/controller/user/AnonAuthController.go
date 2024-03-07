package user

import (
	"context"
	"errors"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/controller/util"
	usecase "server/core/usecase/user"
	"server/infrastructure/logger"

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
	_, err := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアップに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AnonAuthController) SignIn(ctx context.Context, req *connect.Request[user.UserAuthRequest]) (*connect.Response[user.AnonTokenResponse], error) {
	remoteAddr := req.Peer().Addr
	userAgent := req.Header().Get("User-Agent")

	msg := req.Msg
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password, remoteAddr, userAgent)
	if domainErr != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインインに失敗しました。"))
	}
	return connect.NewResponse(&user.AnonTokenResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}

func (ac *AnonAuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[user.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("リセットメールの送信に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
