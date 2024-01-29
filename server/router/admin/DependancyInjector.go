package admin

import (
	usecase "server/core/usecase/admin"

	"server/infrastructure/action"
	implements "server/infrastructure/repository"
)

var (
	authAction                = action.NewAuthClient()
	storeRepository           = implements.NewStoreRepository()
	storeQuery                = implements.NewStoreQueryService()
	adminRepository           = implements.NewAdminRepository()
	adminQuery                = implements.NewAdminQueryService()
	couponRepository          = implements.NewCouponRepository()
	couponQuery               = implements.NewCouponQueryService()
	usercouponQuery           = implements.NewUserCouponQueryService()
	usercouponRepository      = implements.NewUserCouponRepository()
	messageQuery              = implements.NewMessageQueryService()
	messageRepository         = implements.NewMessageRepository()
	mailMagazineQuery         = implements.NewMailMagazineQueryService()
	mailMagazineRepository    = implements.NewMailMagazineRepository()
	mailMagazineLogQuery      = implements.NewMailMagazineLogQueryService()
	mailMagazineLogRepository = implements.NewMailMagazineLogRepository()
	checkinRepository         = implements.NewCheckinRepository()
	userCheckinQuery          = implements.NewCheckinQueryService()
	userDataRepository        = implements.NewUserRepository()
	userDataQuery             = implements.NewUserQueryService()
	sendMailAction            = action.NewSendMailAction()
	fileUploaderAction        = action.NewFileClient()
	userQuery                 = implements.NewUserQueryService()
)

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(adminRepository, adminQuery, storeQuery, authAction)
}

func InitializeAdminDataUsecase() *usecase.AdminDataUsecase {
	return usecase.NewAdminDataUsecase(adminRepository, adminQuery, storeQuery)
}

func InitializeAdminCouponUsecase() *usecase.AdminCouponUsecase {
	return usecase.NewAdminCouponUsecase(couponRepository, couponQuery, usercouponQuery, usercouponRepository, storeQuery)
}

func InitializeUserDataUsecase() *usecase.UserDataUsecase {
	return usecase.NewUserDataUsecase(userDataRepository, userDataQuery)
}

func InitializeMessageUsecase() *usecase.MessageUsecase {
	return usecase.NewMessageUsecase(messageRepository, messageQuery)
}

func InitializeMailMagazineUsecase() *usecase.MailMagazineUsecase {
	return usecase.NewMailMagazineUsecase(mailMagazineRepository, mailMagazineQuery, mailMagazineLogQuery, mailMagazineLogRepository, sendMailAction, userQuery)
}

func InitializeStoreUsecase() *usecase.StoreUsecase {
	return usecase.NewStoreUsecase(storeRepository, storeQuery, fileUploaderAction)
}

func InitializeUserCheckinUsecase() *usecase.UserCheckinUsecase {
	return usecase.NewUserCheckinUsecase(checkinRepository, userCheckinQuery)
}
