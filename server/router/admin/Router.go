package admin

import (
	"net/http"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
)

func RegisterGRPCService(mux *http.ServeMux) *http.ServeMux {
	// リフレクション設定
	reflector := grpcreflect.NewStaticReflector(
		"admin.v1.AdminServer", // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func NewAdminServer(mux *http.ServeMux) {
	// authClient := action.NewAuthClient()
	// requireAuth := router.NewAuthentificatable(authClient)
}
