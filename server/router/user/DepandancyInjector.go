package user

import (
	usecase "server/core/usecase/user"

	action "server/infrastructure/action"
	bookingApiAvail "server/infrastructure/booking/avail"
	bookingApiBook "server/infrastructure/booking/book"
	implements "server/infrastructure/repository"
	"server/infrastructure/wordpress"
)

var (
	userRepo               = implements.NewUserRepository()
	userQueryService       = implements.NewUserQueryService()
	authAction             = action.NewAuthClient()
	userQuery              = implements.NewUserQueryService()
	adminQuery             = implements.NewAdminQueryService()
	storeRepository        = implements.NewStoreRepository()
	checkInRepository      = implements.NewCheckinRepository()
	couponRepository       = implements.NewCouponRepository()
	usercouponRepository   = implements.NewUserCouponRepository()
	usercouponQuery        = implements.NewUserCouponQueryService()
	storeQuery             = implements.NewStoreQueryService()
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
)

// var planQuery = booking.NewPlanQuery()

func InitializeUserUsecase() *usecase.UserDataUsecase {
	return usecase.NewUserDataUsecase(userRepo, userQueryService)
}

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(userRepo, userQueryService, userLoginLogRepository, authAction)
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
	)
}

func InitializeUserReportUsecase() *usecase.UserReportUsecase {
	return usecase.NewUserReportUsecase(
		userQuery,
		userReportRepository,
		sendMailAction,
	)
}
