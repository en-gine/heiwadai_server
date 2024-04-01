package cron

import (
	"net/http"

	cronv1connect "server/api/v1/cron/cronconnect"
	"server/router"

	croncontroller "server/controller/cron"
)

func NewCronServer(mux *http.ServeMux) {
	requireAuthCronHeader := router.NewAuthCronHeader()
	couponUsecase := InitializeCronCouponUsecase()
	couponContoroller := croncontroller.NewCronCouponController(couponUsecase)
	path, handler := cronv1connect.NewCronCouponControllerHandler(couponContoroller, requireAuthCronHeader)
	mux.Handle(path, handler)
}
