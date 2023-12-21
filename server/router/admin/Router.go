package admin

import (
	"net/http"

	adminv1connect "server/api/v1/admin/adminconnect"
	"server/infrastructure/action"
	"server/router"

	admincontroller "server/controller/admin"
)

func NewAdminServer(mux *http.ServeMux) {
	authClient := action.NewAuthClient()
	requireAuth := router.NewAuthentificatable(authClient, userQuery, adminQuery, router.AuthTypeAdmin)

	authUsecase := InitializeAuthUsecase()
	anonAuthContoroller := admincontroller.NewAnonAuthController(authUsecase)
	path, handler := adminv1connect.NewAnonAuthControllerHandler(anonAuthContoroller)
	mux.Handle(path, handler)

	authContoroller := admincontroller.NewAuthController(authUsecase)
	path, handler = adminv1connect.NewAuthControllerHandler(authContoroller, requireAuth)
	mux.Handle(path, handler)

	adminUsecase := InitializeAdminDataUsecase()
	adminContoroller := admincontroller.NewAdminDataController(adminUsecase)
	path, handler = adminv1connect.NewAdminDataControllerHandler(adminContoroller, requireAuth)
	mux.Handle(path, handler)

	couponUsecase := InitializeAdminCouponUsecase()
	couponContoroller := admincontroller.NewAdminCouponController(couponUsecase)
	path, handler = adminv1connect.NewAdminCouponControllerHandler(couponContoroller, requireAuth)
	mux.Handle(path, handler)

	messageUsecase := InitializeMessageUsecase()
	messageContoroller := admincontroller.NewMessageController(messageUsecase)
	path, handler = adminv1connect.NewMessageControllerHandler(messageContoroller, requireAuth)
	mux.Handle(path, handler)

	mailmagazineUsecase := InitializeMailMagazineUsecase()
	mailmagazineContoroller := admincontroller.NewMailMagazineController(mailmagazineUsecase)
	path, handler = adminv1connect.NewMailMagazineControllerHandler(mailmagazineContoroller, requireAuth)
	mux.Handle(path, handler)

	storeUsecase := InitializeStoreUsecase()
	storeContoroller := admincontroller.NewStoreController(storeUsecase)
	path, handler = adminv1connect.NewStoreControllerHandler(storeContoroller, requireAuth)
	mux.Handle(path, handler)

	checkinUsecase := InitializeUserCheckinUsecase()
	checkinContoroller := admincontroller.NewUserCheckinController(checkinUsecase)
	path, handler = adminv1connect.NewUserCheckinControllerHandler(checkinContoroller, requireAuth)
	mux.Handle(path, handler)
}
