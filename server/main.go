package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	adminRouter "server/router/admin"
	userRouter "server/router/user"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	userRouter.NewUserServer(mux)
	adminRouter.NewAdminServer(mux)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), h2c.NewHandler(mux, &http2.Server{}))) // リフレクションを有効にする
}

func InitEnv() {

}
