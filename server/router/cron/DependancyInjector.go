package cron

import (
	usecase "server/core/usecase/cron"

	implements "server/infrastructure/repository"
)

var (
	couponRepository     = implements.NewCouponRepository()
	couponQuery          = implements.NewCouponQueryService()
	usercouponQuery      = implements.NewUserCouponQueryService()
	usercouponRepository = implements.NewUserCouponRepository()
	storeQuery           = implements.NewStoreQueryService()
	transaction          = implements.NewTransaction()
)

func InitializeCronCouponUsecase() *usecase.CronCouponUsecase {
	return usecase.NewCronCouponUsecase(couponRepository, couponQuery, usercouponQuery, usercouponRepository, storeQuery, transaction)
}
