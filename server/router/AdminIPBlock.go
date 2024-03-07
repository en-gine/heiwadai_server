package router

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"server/infrastructure/env"

	"github.com/bufbuild/connect-go"
)

type AdminClientIPFilter struct{}

func NewAdminClientIPFilter() connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			allowIps := env.GetEnv(env.AdminClientIps)

			ip := req.Peer().Addr
			requestIP := strings.Split(ip, ":")[0]

			if !strings.Contains(allowIps, requestIP) {
				fmt.Println("requestIP is blocked: ", requestIP)
				return nil, connect.NewError(connect.CodePermissionDenied, errors.New("IP is invalid"))
			}
			return next(ctx, req)

		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}
