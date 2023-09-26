package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	userRouter "server/router/user"

	"time"
)

var token string

func UserSeeder() {

	authUsecase := userRouter.InitializeAuthUsecase()

	companyName := "株式会社ヒラカワ"
	zip := "810-8861"
	city := "福岡市"
	address := "東区箱崎1-11"
	tel := "0943-77-3185"
	email := os.Getenv("TEST_USER_MAIL")
	fmt.Println(email)
	user, err := authUsecase.Register(
		"Tomohide",
		"Hirakawa",
		"トモヒデ",
		"ヒラカワ",
		&companyName,
		time.Date(1996, 1, 1, 0, 0, 0, 0, time.Local),
		&zip,
		"福岡県",
		&city,
		&address,
		&tel,
		email,
		true,
		true,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if user != nil {
		fmt.Println(user)
	}

	// start server
	http.HandleFunc("/", Handler)
	fmt.Println("Server started on :3000")

	log.Fatal(http.ListenAndServe(":3000", nil))

}
