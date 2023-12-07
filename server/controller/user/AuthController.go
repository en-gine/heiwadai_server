package user

import (
	"context"
	"errors"

	"server/api/v1/user"
	userv1connect "server/api/v1/user/userconnect"
	"server/controller/util"
	usecase "server/core/usecase/user"
	"server/infrastructure/logger"

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
	token, domainErr := ac.authUseCase.SignIn(msg.Email, msg.Password)
	if domainErr != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("サインインに失敗しました。"))
	}
	return connect.NewResponse(&user.UserAuthResponse{
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

func (ac *AuthController) ResetPasswordMail(ctx context.Context, req *connect.Request[user.ResetPasswordRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	err := ac.authUseCase.ResetPasswordMail(msg.Email)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("リセットメールの送信に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (ac *AuthController) UpdatePassword(ctx context.Context, req *connect.Request[user.UpdatePasswordRequest]) (*connect.Response[emptypb.Empty], error) {
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

func (ac *AuthController) UpdateEmail(ctx context.Context, req *connect.Request[user.UpdateEmailRequest]) (*connect.Response[emptypb.Empty], error) {
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

func (ac *AuthController) Refresh(ctx context.Context, req *connect.Request[user.RefreshTokenRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	token := ctx.Value("token").(string)
	if token == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ログインが必要です。"))
	}
	_, err := ac.authUseCase.Refresh(token, msg.RefreshToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("トークンの再取得に失敗しました。"))
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
