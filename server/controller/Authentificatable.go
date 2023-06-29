package controller

import (
	"context"
	"errors"

	"server/core/entity"
	"server/core/infra/action"
	"server/core/infra/types"
	"strings"

	"github.com/bufbuild/connect-go"
)

type Authentificatable struct {
	AuthUser     *entity.User
	Token        string
	RefleshToken string
}

func NewAuthentificatable(AuthClient action.IAuthAction) connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		auth := &Authentificatable{}

		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// リクエストヘッダーからTokenを取得する
			authHeader := req.Header().Get("Authorization")
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: missing token"))
			}
			bearerToken := splitToken[1]

			// リフレッシュトークン取得
			refreshToken := req.Header().Get("X-Refresh-Token")

			// メンバ変数にtokenをセット
			auth.Token = bearerToken
			auth.RefleshToken = refreshToken

			// ユーザーを取得
			user, err := AuthClient.GetUser(&types.Token{AccessToken: auth.Token})

			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			if user == nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			auth.AuthUser = user

			return next(ctx, req)
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
