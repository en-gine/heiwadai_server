package main

import (
	"fmt"
	"log"
	"net/http"

	"server/infrastructure/env"
	adminRouter "server/router/admin"
)

var email = env.GetEnv(env.TestAdminMail)

// password = env.GetEnv(env.TestUserPass)

func AdminSeeder() {
	authUsecase := adminRouter.InitializeAuthUsecase()
	storeUsecase := adminRouter.InitializeStoreUsecase()
	stores, err := storeUsecase.GetList()
	if err != nil {
		panic(err)
	}
	belongStoreID := stores[0].ID

	admin, err := authUsecase.Register(
		"Tomohide",
		belongStoreID,
		email,
	)
	if err != nil {
		panic(err)
	}
	if admin != nil {
		fmt.Println(admin)
	}
	// start server
	http.HandleFunc("/", Handler)
	fmt.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
