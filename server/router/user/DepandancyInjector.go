package user

import (
	usecase "server/core/usecase/user"

	Iaction "server/core/infra/action"
	action "server/infrastructure/action"

	"server/core/infra/repository"

	bookingApiAvail "server/infrastructure/booking/avail"
	bookingApiBook "server/infrastructure/booking/book"
	inmemcache "server/infrastructure/cache"
	"server/infrastructure/env"
	"server/infrastructure/redis"
	implements "server/infrastructure/repository"
	"server/infrastructure/wordpress"
)

var (
	userRepo               = implements.NewUserRepository()
	userQueryService       = implements.NewUserQueryService()
	authAction             = action.NewAuthClient(Iaction.UserTypeUser)
	userQuery              = implements.NewUserQueryService()
	adminQuery             = implements.NewAdminQueryService()
	storeRepository        = implements.NewStoreRepository()
	storeQuery             = implements.NewStoreQueryService()
	checkInRepository      = implements.NewCheckinRepository()
	couponRepository       = implements.NewCouponRepository()
	usercouponRepository   = implements.NewUserCouponRepository()
	usercouponQuery        = implements.NewUserCouponQueryService()
	checkinQuery           = implements.NewCheckinQueryService()
	couponQuery            = implements.NewCouponQueryService()
	transaction            = implements.NewTransaction()
	messageQuery           = implements.NewMessageQueryService()
	messageRepository      = implements.NewMessageRepository()
	bannerQuery            = wordpress.NewBannerQueryService()
	postQuery              = wordpress.NewPostQueryService()
	bookQuery              = implements.NewBookQueryService()
	bookRepository         = implements.NewBookRepository()
	planQuery              = bookingApiAvail.NewPlanQuery(storeQuery)
	bookAPIRepository      = bookingApiBook.NewBookRepository(storeQuery, bookQuery)
	userReportRepository   = implements.NewUserReportRepository()
	userLoginLogRepository = implements.NewUserLoginLogRepository()
	sendMailAction         = action.NewSendMailAction()
	memoryRepository = newMemoryRepository()
)

func newMemoryRepository() repository.IMemoryRepository {
	cacheType := env.GetEnv("CACHE_TYPE")
	if cacheType == "redis" {
		return redis.NewMemoryRepository()
	}
	return inmemcache.NewMemoryRepository()
}

// var planQuery = booking.NewPlanQuery()

func InitializeUserUsecase() *usecase.UserDataUsecase {
	return usecase.NewUserDataUsecase(userRepo, userQueryService)
}

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(userRepo, userQueryService, adminQuery, userLoginLogRepository, authAction)
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
		memoryRepository,
	)
}

func InitializeUserCouponUsecase() *usecase.UserAttachedCouponUsecase {
	return usecase.NewUserAttachedCouponUsecase(
		usercouponRepository,
		usercouponQuery,
		transaction,
	)
}

func InitializeStoreUsecase() *usecase.StoreUsecase {
	return usecase.NewStoreUsecase(
		storeQuery,
	)
}

func InitializePlanUsecase() *usecase.PlanUsecase {
	return usecase.NewPlanUsecase(
		planQuery,
		storeQuery,
	)
}

func InitializeMessageUsecase() *usecase.MessageUsecase {
	return usecase.NewMessageUsecase(
		messageRepository,
		messageQuery,
	)
}

func InitializeBookUsecase() *usecase.BookUsecase {
	return usecase.NewBookUsecase(
		bookQuery,
		bookRepository,
		bookAPIRepository,
		sendMailAction,
		storeQuery,
	)
}

func InitializeUserReportUsecase() *usecase.UserReportUsecase {
	return usecase.NewUserReportUsecase(
		userQuery,
		userReportRepository,
		sendMailAction,
	)
}
