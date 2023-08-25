package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	userv1connect "server/api/v1/user/userconnect"
	"server/infrastructure/action"

	usercontroller "server/controller/user"
	"server/router"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewConnectReflection() *http.ServeMux {
	// リフレクション設定
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserServer", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func NewUserServer() {

	mux := NewConnectReflection()
	authClient := action.NewAuthClient()
	requireAuth := router.NewAuthentificatable(&authClient)

	authUsecase := InitializeAuthUsecase()
	authContoroller := usercontroller.NewAuthController(authUsecase)
	path, handler := userv1connect.NewAuthControllerHandler(authContoroller)
	mux.Handle(path, handler)

	userUsecase := InitializeUserUsecase()
	userContoroller := usercontroller.NewUserDataController(userUsecase)
	path, handler = userv1connect.NewUserDataControllerHandler(userContoroller, requireAuth)
	mux.Handle(path, handler)

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), h2c.NewHandler(mux, &http2.Server{}))) // リフレクションを有効にする

}
