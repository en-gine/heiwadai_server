package router

import (
	"context"
	"errors"
	"strconv"

	"server/core/infra/action"
	"strings"

	"github.com/bufbuild/connect-go"
)

type keyType string

var UserIDKey keyType = "userID"

type Authentificatable struct {
}

func NewAuthentificatable(AuthClient action.IAuthAction) connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		// auth := &Authentificatable{}

		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// リクエストヘッダーからTokenを取得する
			authHeader := req.Header().Get("Authorization")
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: トークンが見つかりません。"))
			}
			bearerToken := splitToken[1]

			// リフレッシュトークン取得
			refreshToken := req.Header().Get("X-Refresh-Token")

			Token, err := AuthClient.Refresh(bearerToken, refreshToken)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, errors.New("error: リフレッシュトークンの取得に問題が発生しました。"))
			}

			id, error := AuthClient.GetUserID(bearerToken)
			if error != nil {
				return nil, connect.NewError(connect.CodeInternal, errors.New("error: ユーザーの取得に問題が発生しました。"))
			}
			if id == nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: 正しいトークンでは無いか有効期限が切れています。"))
			}
			ctx = context.WithValue(ctx, UserIDKey, id)

			res, err := next(ctx, req)
			if Token != nil {
				res.Header().Set("AccessToken:", Token.AccessToken)
				res.Header().Set("RefreshToken:", *Token.RefreshToken)
				res.Header().Set("Expire:", strconv.Itoa(*Token.ExpiresIn))
			}
			return res, err

		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
