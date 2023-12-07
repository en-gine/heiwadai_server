package admin

import (
	"context"
	"errors"

	"server/api/v1/admin"
	adminv1connect "server/api/v1/admin/adminconnect"
	usecase "server/core/usecase/admin"
	"server/infrastructure/logger"

	connect "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	authUseCase usecase.AuthUsecase
}

var _ adminv1connect.AuthControllerClient = &AuthController{}

func NewAuthController(authUsecase *usecase.AuthUsecase) *AuthController {
	return &AuthController{
		authUseCase: *authUsecase,
	}
}

func (ac *AuthController) Register(ctx context.Context, req *connect.Request[admin.AdminRegisterRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg

	_, domaiErr := ac.authUseCase.Register(
		msg.Name,
		uuid.MustParse(msg.BelongStoreID),
		msg.Mail,
	)
	return connect.NewResponse(&emptypb.Empty{}), domaiErr
}

func (ac *AuthController) SignUp(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.SignUp(msg.Email, msg.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアップに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) SignIn(ctx context.Context, req *connect.Request[admin.AdminAuthRequest]) (*connect.Response[admin.AdminAuthResponse], error) {
	msg := req.Msg
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password)
	if domainErr != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインインに失敗しました。"))
	}
	return connect.NewResponse(&admin.AdminAuthResponse{
		AccessToken:  token.AccessToken,
		ExpiresIn:    int64(*token.ExpiresIn),
		RefreshToken: *token.RefreshToken,
	}), nil
}

func (ac *AuthController) SignOut(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	token := ctx.Value("token").(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.SignOut(token)
	if err != nil {
		logger.Error(err.Error())
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインアウトに失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[admin.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("リセットメールの送信に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect.Request[admin.UpdatePasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	token := ctx.Value("token").(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.UpdatePassword(msg.Password, token)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("パスワードの変更に失敗しました。\nネットワークの問題や同じパスワードに変更した、などの理由が考えられます。"))
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect.Request[admin.UpdateEmailRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	token := ctx.Value("token").(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	err := ac.authUseCase.UpdateEmail(msg.Email, token)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("メールアドレスの変更に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) Refresh(ctx context.Context, req *connect.Request[admin.AdminRefreshTokenRequest]) (*connect.Response[admin.AdminAuthResponse], error) {
	msg := req.Msg
	token := ctx.Value("token").(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	tkn, err := ac.authUseCase.Refresh(token, msg.RefreshToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("トークンの再取得に失敗しました。"))
	}
	return connect.NewResponse(&admin.AdminAuthResponse{
		AccessToken:  tkn.AccessToken,
		ExpiresIn:    int64(*tkn.ExpiresIn),
		RefreshToken: *tkn.RefreshToken,
	}), nil
}
