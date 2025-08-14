package user

import (
	"context"
	"errors"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller"
	"server/controller/util"
	domainErrors "server/core/errors"
	usecase "server/core/usecase/user"
	"server/router"

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

func (ac *AuthController) SignOut(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	if ctx.Value(router.TokenKey) == nil {
		return connect.NewResponse(&emptypb.Empty{}), nil
	}
	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return connect.NewResponse(&emptypb.Empty{}), nil
	}
	domainErr := ac.authUseCase.SignOut(token)
	if domainErr != nil {
		// セッションが既に無効の場合もサインアウト成功として扱う
		if domainErr.GetType() == domainErrors.QueryDataNotFoundError {
			return connect.NewResponse(&emptypb.Empty{}), nil
		}
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect.Request[user.UpdatePasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	if ctx.Value(router.TokenKey) == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("トークンが必要です。"))
	}
	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}

	domainErr := ac.authUseCase.UpdatePassword(msg.Password, token)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect.Request[user.UpdateEmailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	if ctx.Value(router.TokenKey) == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("トークンが必要です。"))
	}
	token := ctx.Value(router.TokenKey).(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	domainErr := ac.authUseCase.UpdateEmail(msg.Email, token)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) Refresh(ctx context.Context, req *connect.Request[user.RefreshTokenRequest]) (*connect.Response[user.UserAuthTokenResponse], error) {
	msg := req.Msg
	token := msg.AccessToken
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	if msg.RefreshToken == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("リフレッシュトークンが必要です。"))
	}

	tkn, domainErr := ac.authUseCase.Refresh(token, msg.RefreshToken)
	if domainErr != nil {
		return nil, controller.ErrorHandler(domainErr)
	}
	return connect.NewResponse(&user.UserAuthTokenResponse{
		AccessToken:  tkn.AccessToken,
		ExpiresIn:    int64(*tkn.ExpiresIn),
		RefreshToken: *tkn.RefreshToken,
	}), nil
}
