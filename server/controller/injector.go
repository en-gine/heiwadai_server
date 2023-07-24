package controller

import (
	usecase "server/core/usecase/user"

	action "server/infrastructure/action"
	implements "server/infrastructure/repository"
)

var userRepo = implements.NewUserRepository()
var userQueryService = implements.NewUserQueryService()
var authAction = action.NewAuthClient()

func InitializeUserUsecase() *usecase.UserDataUsecase {
	return usecase.NewUserDataUsecase(userRepo, userQueryService)
}

func InitializeAuthUsecase() *usecase.AuthUsecase {
	return usecase.NewAuthUsecase(userRepo, userQueryService, &authAction)
}
