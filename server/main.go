package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	userv1connect "server/api/v1/user/userconnect"

	controller "server/controller/user"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var Logger *zap.Logger

// リフレクション設定
func newServeMuxWithReflection() *http.ServeMux {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserServer", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

// インターセプタ設定
func newInterCeptors() connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// ここでヘッダをセットするなど色々処理を書ける
			Logger.Info(
				"Request",
				zap.String("Procedure", req.Spec().Procedure),
				zap.String("Protocol", req.Peer().Protocol),
				zap.String("Addr", req.Peer().Addr),
			)
			return next(ctx, req)
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}

func main() {

	userContoroller := controller.UserRegisterController{}
	mux := newServeMuxWithReflection()
	interceptor := newInterCeptors()
	path, handler := userv1connect.NewUserRegisterControllerHandler(&userContoroller, interceptor)
	mux.Handle(path, handler)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), h2c.NewHandler(mux, &http2.Server{}))) // Use h2c so we can serve HTTP/2 without TLS.

}
