package main

import (
	"fmt"
	"server/router/user"

	"github.com/google/uuid"
)

func main() {
	StampCard()
	// Checkin()
}

func Checkin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	limitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	unlimitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	fmt.Print(unlimitQr)
	coupon, err := usecase.Checkin(userID, limitQr)
	if err != nil {
		panic(err)
	}
	println(coupon)
}

func StampCard() {
	usecase := user.InitializeUserCheckinUsecase()
	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	card, err := usecase.GetStampCard(userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(card.Checkins[0])

}
