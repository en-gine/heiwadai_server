package router

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/env"
	"server/infrastructure/redis"

	"github.com/bufbuild/connect-go"
)

type keyType string

type AuthType string

var (
	AuthTypeAdmin AuthType = "admin"
	AuthTypeUser  AuthType = "user"
)

var (
	UserIDKey    keyType = "userID"
	UserAuthType keyType = "authType"
	TokenKey     keyType = "token"
)

var cache = redis.NewMemoryRepository()

type Authentificatable struct{}

func NewAuthentificatable(AuthClient action.IAuthAction, UserDataQuery queryservice.IUserQueryService, AdminDataQuery queryservice.IAdminQueryService, AuthType AuthType) connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		// auth := &Authentificatable{}

		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// リクエストヘッダーからTokenを取得する
			authHeader := req.Header().Get("Authorization")
			if authHeader == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証ヘッダーが必要なアクセスです。"))
			}

			bearerToken := strings.Replace(authHeader, "Bearer ", "", 1)
			if bearerToken == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("トークンが見つかりません。"))
			}

			var authData *action.UserAuth
			// キャッシュから取得
			byteCache := cache.Get(bearerToken)

			if byteCache != nil {
				err := json.Unmarshal(*byteCache, &authData)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, errors.New("キャッシュの取得に問題が発生しました。"))
				}

			} else { // 再取得API
				var err error // err変数の宣言
				// リフレッシュトークン取得
				refreshToken := req.Header().Get("X-Refresh-Token")

				authData, err = AuthClient.Refresh(bearerToken, refreshToken)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, errors.New("リフレッシュトークンの取得に問題が発生しました。"))
				}

				authJSON, err := json.Marshal(authData)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, errors.New("キャッシュの保存に問題が発生しました。"))
				}

				// キャッシュに保存
				cache.Set(bearerToken, authJSON, time.Duration(*authData.Token.ExpiresIn)*time.Second)
			}

			userID := authData.UserID
			token := authData.Token
			userType := authData.UserType

			if AuthType == AuthTypeAdmin && userType.String() != "admin" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("管理者として登録されていません。"))
			}

			if AuthType == AuthTypeUser && userType.String() != "user" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("ユーザーとして登録されていません。"))
			}

			if token != nil {
				ctx = context.WithValue(ctx, TokenKey, token.AccessToken)
			}
			ctx = context.WithValue(ctx, UserIDKey, userID.String())
			if env.GetEnv(env.EnvMode) == "dev" {
				fmt.Println("----------------reqest----------------")
				fmt.Println(req)
				fmt.Println("----------------userID--------------")
				fmt.Println(userID)
			}
			res, err := next(ctx, req)
			if err != nil {
				// logger.Error(err.Error())
				return nil, err
			}
			if env.GetEnv(env.EnvMode) == "dev" {
				fmt.Println("----------------response--------------")
				fmt.Println(res)
			}

			if token != nil && res != nil {
				res.Header().Set("AccessToken", token.AccessToken)
				res.Header().Set("RefreshToken", *token.RefreshToken)
				res.Header().Set("ExpiresIn", strconv.Itoa(*token.ExpiresIn))
			}
			return res, nil
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
