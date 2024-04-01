package router

import (
	"context"
	"fmt"

	"server/infrastructure/env"

	"github.com/bufbuild/connect-go"
)

func NewLogger() connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if env.GetEnv(env.EnvMode) == "dev" {
				fmt.Println("----------------reqest----------------")
				fmt.Println(req.Header())
			}
			res, err := next(ctx, req)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if env.GetEnv(env.EnvMode) == "dev" {
				fmt.Println("----------------response--------------")
				fmt.Println(res)
			}

			return res, nil
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
