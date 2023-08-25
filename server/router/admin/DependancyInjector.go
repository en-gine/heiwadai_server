package admin

import (
	usecase "server/core/usecase/admin"

	implements "server/infrastructure/repository"
)

var storeRepo = implements.NewStoreRepository()
var storeQueryService = implements.NewStoreQueryService()

func InitializeStoreUsecase() *usecase.StoreUsecase {
	return usecase.NewStoreUsecase(storeRepo, storeQueryService)
}
