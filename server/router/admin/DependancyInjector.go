package admin

import (
	usecase "server/core/usecase/admin"

	"server/infrastructure/action"
	implements "server/infrastructure/repository"
)

var authAction = action.NewAuthClient()
var storeRepo = implements.NewStoreRepository()
var storeQueryService = implements.NewStoreQueryService()
var adminRepo = implements.NewAdminRepository()
var adminQueryService = implements.NewAdminQueryService()

func InitializeStoreUsecase() *usecase.StoreUsecase {
	return usecase.NewStoreUsecase(storeRepo, storeQueryService)
}

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(adminRepo, adminQueryService, storeQueryService, authAction)
}
