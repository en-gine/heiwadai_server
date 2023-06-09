package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	userv1 "server/api/user/v1"
	"server/api/user/v1/userv1connect"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UserServer struct{}

func (s *UserServer) User(ctx context.Context, req *connect.Request[userv1.UserRequest]) (*connect.Response[userv1.UserResponse], error) {
	log.Println("Request headers: ", req.Header())

	if req.Msg.Name == "" {
		// エラーにステータスコードを追加
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("name is required."))
	}

	userResp := &userv1.UserResponse{
		Usering: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	}
	resp := connect.NewResponse(userResp)
	// ヘッダをセットしてみたり
	resp.Header().Set("Greet-Version", "v1")
	return resp, nil
}

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
			req.Header().Set("hoge", "fuga")
			return next(ctx, req)
		})
	}
	return connect.WithInterceptors(connect.UnaryInterceptorFunc(interceptor))
}

func main() {

	userServer := &UserServer{}
	mux := newServeMuxWithReflection()
	interceptor := newInterCeptors()
	path, handler := userv1connect.NewUserServiceHandler(userServer, interceptor)
	mux.Handle(path, handler)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), h2c.NewHandler(mux, &http2.Server{}))) // Use h2c so we can serve HTTP/2 without TLS.

}
