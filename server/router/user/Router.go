package user

import (
	"net/http"

	userv1connect "server/api/v1/user/userconnect"
	"server/infrastructure/action"

	usercontroller "server/controller/user"
	"server/router"
)

func NewUserServer(mux *http.ServeMux) {
	authClient := action.NewAuthClient()
	requireAuth := router.NewAuthentificatable(authClient, userQuery, adminQuery, router.AuthTypeUser)

	authUsecase := InitializeAuthUsecase()
	anonAuthContoroller := usercontroller.NewAnonAuthController(authUsecase)
	path, handler := userv1connect.NewAnonAuthControllerHandler(anonAuthContoroller)
	mux.Handle(path, handler)

	authContoroller := usercontroller.NewAuthController(authUsecase)
	path, handler = userv1connect.NewAuthControllerHandler(authContoroller, requireAuth)
	mux.Handle(path, handler)

	userUsecase := InitializeUserUsecase()
	userContoroller := usercontroller.NewUserDataController(userUsecase)
	path, handler = userv1connect.NewUserDataControllerHandler(userContoroller, requireAuth)
	mux.Handle(path, handler)

	bannerUsecase := InitializeBannerUsecase()
	bannerContoroller := usercontroller.NewBannerController(bannerUsecase)
	path, handler = userv1connect.NewBannerControllerHandler(bannerContoroller, requireAuth)
	mux.Handle(path, handler)

	postUsecase := InitializePostUsecase()
	postContoroller := usercontroller.NewPostController(postUsecase)
	path, handler = userv1connect.NewPostControllerHandler(postContoroller, requireAuth)
	mux.Handle(path, handler)

	storeUsecase := InitializeStoreUsecase()
	storeContoroller := usercontroller.NewStoreController(storeUsecase)
	path, handler = userv1connect.NewStoreControllerHandler(storeContoroller, requireAuth)
	mux.Handle(path, handler)

	couponUsecase := InitializeUserCouponUsecase()
	couponContoroller := usercontroller.NewMyCouponController(couponUsecase)
	path, handler = userv1connect.NewMyCouponControllerHandler(couponContoroller, requireAuth)
	mux.Handle(path, handler)

	checkinUsecase := InitializeUserCheckinUsecase()
	checkinContoroller := usercontroller.NewCheckInController(checkinUsecase)
	path, handler = userv1connect.NewCheckinControllerHandler(checkinContoroller, requireAuth)
	mux.Handle(path, handler)

	bookUsecase := InitializeBookUsecase()
	bookContoroller := usercontroller.NewBookController(bookUsecase)
	path, handler = userv1connect.NewBookControllerHandler(bookContoroller, requireAuth)
	mux.Handle(path, handler)

	messageUsecase := InitializeMessageUsecase()
	messageContoroller := usercontroller.NewMessageController(messageUsecase)
	path, handler = userv1connect.NewMessageControllerHandler(messageContoroller, requireAuth)
	mux.Handle(path, handler)

	userReportUsecase := InitializeUserReportUsecase()
	userReportContoroller := usercontroller.NewUserReportController(userReportUsecase)
	path, handler = userv1connect.NewUserReportControllerHandler(userReportContoroller, requireAuth)
	mux.Handle(path, handler)
}
