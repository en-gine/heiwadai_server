package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	adminRouter "server/router/admin"
)

func AdminSeeder() {
	authUsecase := adminRouter.InitializeAuthUsecase()
	storeUsecase := adminRouter.InitializeStoreUsecase()
	stores, err := storeUsecase.GetList()
	if err != nil {
		panic(err)
	}
	belongStoreID := stores[0].ID

	email := os.Getenv("TEST_ADMIN_MAIL")

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
