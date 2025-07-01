package router

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	domainErrors "server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	inmemcache "server/infrastructure/cache"
	"server/infrastructure/env"
	"server/infrastructure/logger"
	"server/infrastructure/redis"
	"strconv"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
)

type keyType string

type AuthType string

var (
	AuthTypeAdmin AuthType = "admin"
	AuthTypeUser  AuthType = "user"
)

func (ut AuthType) String() string {
	return string(ut)
}

var (
	UserIDKey    keyType = "userID"
	UserAuthType keyType = "authType"
	TokenKey     keyType = "token"
)

var cache repository.IMemoryRepository

type Authentificatable struct{}

func init() {
	// 環境変数でキャッシュタイプを切り替え
	cacheType := env.GetEnv("CACHE_TYPE")
	if cacheType == "" {
		cacheType = "memory" // デフォルトはインメモリ
	}
	
	switch cacheType {
	case "redis":
		cache = redis.NewMemoryRepository()
	case "memory":
		cache = inmemcache.NewMemoryRepository()
	default:
		cache = inmemcache.NewMemoryRepository()
	}
}

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

				var domainErr *domainErrors.DomainError
				authData, domainErr, err = AuthClient.Refresh(bearerToken, refreshToken)
				if domainErr != nil {
					return nil, domainErr
				}

				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, errors.New("リフレッシュトークンの取得に問題が発生しました。"))
				}

				authJSON, err := json.Marshal(authData)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, errors.New("キャッシュの保存に問題が発生しました。"))
				}

				if authData == nil {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証情報が取得できませんでした。"))
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
				logger.DebugPrint(req)
			}
			res, err := next(ctx, req)
			if err != nil {
				// logger.Error(err.Error())
				return nil, err
			}
			if env.GetEnv(env.EnvMode) == "dev" {
				fmt.Println("----------------response--------------")
				logger.DebugPrint(res)
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
