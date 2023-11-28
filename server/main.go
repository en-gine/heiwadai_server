package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"server/infrastructure/env"
	adminRouter "server/router/admin"
	userRouter "server/router/user"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	env.InitEnv() // 環境変数を読み込む

	mux := http.NewServeMux()

	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "heiwadai app server!")
	})

	msg := os.ExpandEnv("${ENV_MODE} mode run! port: ${PORT}")
	fmt.Println(msg)
	EchoMyIP()
	port := env.GetEnv(env.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{})))) // リフレクションを有効にする
}
