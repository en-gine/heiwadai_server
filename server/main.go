package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"server/infrastructure/env"
	"server/infrastructure/logger"
	"server/infrastructure/redis"
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
	os.Setenv("TZ", "Asia/Tokyo")

	rdb := redis.NewMemoryRepository()             // redis接続
	rdb.Set("key", []byte("test"), time.Second*10) // redisにデータを保存
	data := rdb.Get("key")                         // redisからデータを取得
	if data == nil || string(*data) != "test" {
		logger.Warn("redisにはデータが保存されません。")
	}

	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)
	RegisterGRPCService(mux) // リフレクションを有効にする

	msg := os.ExpandEnv("${ENV_MODE} mode run! port: ${PORT}")
	fmt.Println(msg)
	EchoMyIP()
	port := env.GetEnv(env.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, AllowCors().Handler(h2c.NewHandler(mux, &http2.Server{})))) // リフレクションを有効にする
}

func AllowCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Authorization",
			"X-Cron-Key",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Grpc-Timeout",
			"X-Grpc-Web",
			"X-User-Agent",
			"X-Refresh-Token",
			"User-Agent",
			"Referer",
			"Connection",
			"Cookie",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
			"Origin",
		},
		AllowCredentials: true,
	})
}

func RegisterGRPCService(mux *http.ServeMux) *http.ServeMux {
	// リフレクション設定
	reflector := grpcreflect.NewStaticReflector(
		"server.user.AnonAuthController",
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
		"server.admin.AnonAuthController",
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
