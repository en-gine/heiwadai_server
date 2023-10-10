package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/infrastructure/env"
	adminRouter "server/router/admin"
	userRouter "server/router/user"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	env.InitEnv() // 環境変数を読み込む

	mux := http.NewServeMux()
	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)
	port := env.GetEnv(env.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, h2c.NewHandler(mux, &http2.Server{}))) // リフレクションを有効にする
}
