package main

import (
	"server/router/user"

	"github.com/google/uuid"
)

func main() {
	// Search()
	Checkin()
}

func Checkin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	storeID, _ := uuid.Parse("2163a9de-b196-420a-a429-21c88deade23")
	coupon, err := usecase.Checkin(userID, storeID)
	if err != nil {
		panic(err)
	}
	println(coupon)

}
