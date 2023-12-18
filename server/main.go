package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/infrastructure/env"
	adminRouter "server/router/admin"
	userRouter "server/router/user"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	env.InitEnv() // 環境変数を読み込む

	mux := http.NewServeMux()

	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)
	RegisterGRPCService(mux)

	msg := os.ExpandEnv("${ENV_MODE} mode run! port: ${PORT}")
	fmt.Println(msg)
	EchoMyIP()
	port := env.GetEnv(env.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{})))) // リフレクションを有効にする
}

func RegisterGRPCService(mux *http.ServeMux) *http.ServeMux {
	// リフレクション設定
	reflector := grpcreflect.NewStaticReflector(
		"server.user.AuthController",
		"server.user.BannerController",
		"server.user.BookController",
		"server.user.CheckinController",
		"server.user.MessageController",
		"server.user.MyCouponController",
		"server.user.PostController",
		"server.user.StoreController",
		"server.user.UserDataController",
		"server.user.UserReportController",
		"server.admin.AuthController",
		"server.admin.AdminDataController",
		"server.admin.AuthController",
		"server.admin.AdminCouponController",
		"server.admin.MailMagazineController",
		"server.admin.MessageController",
		"server.admin.StoreController",
		"server.admin.UserCheckinController",
		"server.admin.UserDataController",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}
