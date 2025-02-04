package router

import (
	"context"
	"errors"

	"server/infrastructure/env"
	"server/infrastructure/logger"

	"github.com/bufbuild/connect-go"
)

type AuthCronHeader struct{}

func NewAuthCronHeader() connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// リクエストヘッダーからTokenを取得する
			cronSecret := req.Header().Get("Authorization")
			cronKey := req.Header().Get("X-Cron-Key")
			logger.DebugPrint(req.Header())

			if cronSecret == "" || cronKey == "" {
				logger.Error("認証ヘッダーが必要なアクセスです。")
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証ヘッダーが必要なアクセスです。"))
			}

			if cronSecret != env.GetEnv(env.CronAccessSecret) {
				logger.Error("認証情報（シークレットキー）が正しくありません。")
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証情報（シークレットキー）が正しくありません。"))
			}
			if cronKey != env.GetEnv(env.CronAccessKey) {
				logger.Error("認証情報（キー情報）が正しくありません。")
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証情報（キー情報）が正しくありません。"))
			}

			// allowIps := env.GetEnv(env.AdminClientIps)

			// ip := req.Peer().Addr
			// requestIP := strings.Split(ip, ":")[0]

			// if !strings.Contains(allowIps, requestIP) {
			// 	fmt.Println("requestIP is blocked: ", requestIP)
			// 	return nil, connect.NewError(connect.CodePermissionDenied, errors.New("IP is invalid"))
			// }

			return next(ctx, req)
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
