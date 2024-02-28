package cron

import (
	usecase "server/core/usecase/cron"

	"server/infrastructure/redis"
	implements "server/infrastructure/repository"
)

var (
	couponRepository     = implements.NewCouponRepository()
	couponQuery          = implements.NewCouponQueryService()
	usercouponQuery      = implements.NewUserCouponQueryService()
	usercouponRepository = implements.NewUserCouponRepository()
	userDataRepository   = implements.NewUserRepository()
	userDataQuery        = implements.NewUserQueryService()
	userQuery            = implements.NewUserQueryService()
	storeQuery           = implements.NewStoreQueryService()
	transaction          = implements.NewTransaction()
	cache                = redis.NewMemoryRepository()
)

func InitializeCronCouponUsecase() *usecase.CronCouponUsecase {
	return usecase.NewCronCouponUsecase(couponRepository, couponQuery, usercouponQuery, usercouponRepository, storeQuery, transaction)
}
