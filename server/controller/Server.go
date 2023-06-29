package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	userv1connect "server/api/v1/user/userconnect"
	"server/external/spabase"

	controller "server/controller/user"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var Logger *zap.Logger

// リフレクション設定
func NewConnectReflection() *http.ServeMux {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserServer", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func StartServer() {

	mux := NewConnectReflection()
	requireAuth := NewAuthentificatable(spabase.NewAuthClient())
	userContoroller := controller.UserRegisterController{}
	path, handler := userv1connect.NewUserRegisterControllerHandler(&userContoroller, requireAuth)
	mux.Handle(path, handler)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), h2c.NewHandler(mux, &http2.Server{}))) // リフレクションを有効にする

}
