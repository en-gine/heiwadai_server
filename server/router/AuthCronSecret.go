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
			if cronSecret == "" || cronKey == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証ヘッダーが必要なアクセスです。"))
			}

			if cronSecret != env.GetEnv(env.CronAccessSecret) {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証情報が正しくありません。"))
			}
			if cronKey != env.GetEnv(env.CronAccessKey) {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("認証情報が正しくありません。"))
			}

			logger.Infof("cron request: %s", req)
			res, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			logger.Infof("cron response: %s", req)

			return res, nil
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
