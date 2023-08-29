package user

import (
	"net/http"

	userv1connect "server/api/v1/user/userconnect"
	"server/infrastructure/action"

	usercontroller "server/controller/user"
	"server/router"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
)

func RegisterGRPCService(mux *http.ServeMux) *http.ServeMux {
	// リフレクション設定
	reflector := grpcreflect.NewStaticReflector(
		"user.v1.UserServer", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func NewUserServer(mux *http.ServeMux) {

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

}
