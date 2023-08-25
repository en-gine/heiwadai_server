package main

import (
	usecaseAdmin "server/core/usecase/admin"
	usecaseUser "server/core/usecase/user"

	action "server/infrastructure/action"
	implements "server/infrastructure/repository"
)

var userRepo = implements.NewUserRepository()
var userQueryService = implements.NewUserQueryService()
var authAction = action.NewAuthClient()
var storeRepo = implements.NewStoreRepository()
var storeQueryService = implements.NewStoreQueryService()

func InitializeUserUsecase() *usecaseUser.UserDataUsecase {
	return usecaseUser.NewUserDataUsecase(userRepo, userQueryService)
}

func InitializeAuthUsecase() *usecaseUser.AuthUsecase {
	return usecaseUser.NewAuthUsecase(userRepo, userQueryService, &authAction)
}

func InitializeStoreUsecase() *usecaseAdmin.StoreUsecase {
	return usecaseAdmin.NewStoreUsecase(storeRepo, storeQueryService)
}
