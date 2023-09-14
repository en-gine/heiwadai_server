package user

import (
	usecase "server/core/usecase/user"

	action "server/infrastructure/action"
	"server/infrastructure/booking"
	implements "server/infrastructure/repository"
	"server/infrastructure/wordpress"
)

var userRepo = implements.NewUserRepository()
var userQueryService = implements.NewUserQueryService()
var authAction = action.NewAuthClient()
var userQuery = implements.NewUserQueryService()
var storeRepository = implements.NewStoreRepository()
var checkInRepository = implements.NewCheckinRepository()
var couponRepository = implements.NewCouponRepository()
var usercouponRepository = implements.NewUserCouponRepository()
var usercouponQuery = implements.NewUserCouponQueryService()
var storeQuery = implements.NewStoreQueryService()
var checkinQuery = implements.NewCheckinQueryService()
var couponQuery = implements.NewCouponQueryService()
var transaction = implements.NewTransaction()
var bannerQuery = wordpress.NewBannerQueryService()
var postQuery = wordpress.NewPostQueryService()
var planQuery = booking.NewPlanQuery()

func InitializeUserUsecase() *usecase.UserDataUsecase {
	return usecase.NewUserDataUsecase(userRepo, userQueryService)
}

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(userRepo, userQueryService, authAction)
}

func InitializeBannerUsecase() *usecase.BannerUsecase {
	return usecase.NewBannerUsecase(bannerQuery)
}

func InitializePostUsecase() *usecase.PostUsecase {
	return usecase.NewPostUsecase(postQuery)
}

func InitializeUserCheckinUsecase() *usecase.UserCheckinUsecase {
	return usecase.NewUserCheckinUsecase(
		userQuery,
		storeRepository,
		checkInRepository,
		couponRepository,
		usercouponRepository,
		usercouponQuery,
		storeQuery,
		checkinQuery,
		couponQuery,
		transaction,
	)
}

func InitializeUserCouponUsecase() *usecase.UserAttachedCouponUsecase {
	return usecase.NewUserAttachedCouponUsecase(
		usercouponRepository,
		usercouponQuery,
	)
}

func InitializeStoreUsecase() *usecase.StoreUsecase {
	return usecase.NewStoreUsecase(
		storeQuery,
	)
}

// func InitializePlanUsecase() *usecase.PlanUsecase {
// 	return usecase.NewPlanUsecase(
// 		planQuery,
// 	)
// }
