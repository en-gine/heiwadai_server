package user

import (
	usecase "server/core/usecase/user"

	action "server/infrastructure/action"
	implements "server/infrastructure/repository"
	"server/infrastructure/wordpress"
)

var (
	userRepo             = implements.NewUserRepository()
	userQueryService     = implements.NewUserQueryService()
	authAction           = action.NewAuthClient()
	userQuery            = implements.NewUserQueryService()
	storeRepository      = implements.NewStoreRepository()
	checkInRepository    = implements.NewCheckinRepository()
	couponRepository     = implements.NewCouponRepository()
	usercouponRepository = implements.NewUserCouponRepository()
	usercouponQuery      = implements.NewUserCouponQueryService()
	storeQuery           = implements.NewStoreQueryService()
	checkinQuery         = implements.NewCheckinQueryService()
	couponQuery          = implements.NewCouponQueryService()
	transaction          = implements.NewTransaction()
	bannerQuery          = wordpress.NewBannerQueryService()
	postQuery            = wordpress.NewPostQueryService()
)

// var planQuery = booking.NewPlanQuery()

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
