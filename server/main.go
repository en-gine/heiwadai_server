package main

//go:generate go run linter/main.go .

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server/infrastructure/env"
	"server/infrastructure/redis"
	"server/infrastructure/repository"
	adminRouter "server/router/admin"
	cronRouter "server/router/cron"
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
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = location
	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)
	cronRouter.NewCronServer(mux)
	RegisterGRPCService(mux) // リフレクションを有効にする
	mux.HandleFunc("/", Index)

	fmt.Println(os.ExpandEnv("${ENV_MODE} mode run! port: ${PORT}"))
	fmt.Println(CheckMyIP())
	fmt.Println(CheckRedisStatus())
	fmt.Println(CheckDBStatus())
	fmt.Println("サーバー時刻:", time.Now())
	port := env.GetEnv(env.ServerPort)
	
	// Setup graceful shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: AllowCors().Handler(h2c.NewHandler(mux, &http2.Server{})),
	}
	
	// Start server in a goroutine
	go func() {
		// adminに対する各種IP制限はインターセプタの中で行っています。
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	
	// Wait for interrupt signal to gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	
	fmt.Println("\nShutting down server...")
	
	// Shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// Shutdown HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
	
	// Close database connection
	// InitDB() returns the connection, so we need to manage it properly
	fmt.Println("Database connection will be closed by connection pool timeout")
	
	fmt.Println("Server gracefully stopped")
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
			"X-Real-IP",
			"X-Forwarded-For",
			"User-Agent",
			"Host",
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
		"server.cron.CronCouponController",
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Heiwadai Server is running! :)")
	fmt.Fprintln(w, CheckMyIP())
	fmt.Fprintln(w, CheckRedisStatus())
	fmt.Fprintln(w, CheckDBStatus())
	fmt.Fprintln(w, "サーバー時刻:", time.Now())
}

func CheckRedisStatus() string {
	rdb := redis.NewMemoryRepository()             // redis接続
	rdb.Set("key", []byte("test"), time.Second*10) // redisにデータを保存
	data := rdb.Get("key")                         // redisからデータを取得
	if data == nil || string(*data) != "test" {
		return "redis connection error :("
	}
	return "redis connection success!"
}

func CheckDBStatus() string {
	db := repository.InitDB()
	if db == nil {
		return "DB connection error :("
	}
	err := db.Ping()
	if err != nil {
		return "DB connection error :( \n" + err.Error()
	}
	return "DB connection success!"
}
