package main

import (
	userRouter "server/router/user"

	"time"
)

func UserSeeder() {
	authUsecase := userRouter.InitializeAuthUsecase()

	companyName := "株式会社ヒラカワ"
	zip := "810-8861"
	city := "福岡市"
	address := "東区箱崎1-11"
	tel := "0943-77-3185"

	authUsecase.Register(
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
		"sutefu23@gmail.com",
		true,
		true,
	)
}
