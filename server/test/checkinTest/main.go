package main

import (
	"fmt"

	"server/infrastructure/logger"
	"server/router/user"

	"github.com/google/uuid"
)

func main() {
	// StampCard()
	UnlimitCheckin()
	// Checkin()
}

func Checkin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	limitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	coupon, err := usecase.Checkin(userID, limitQr)
	if err != nil {
		logger.Error(err.Error())
	}
	println(coupon)
}

func UnlimitCheckin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	unlimitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	coupon, err := usecase.Checkin(userID, unlimitQr)
	if err != nil {
		logger.Error(err.Error())
	}
	println(coupon)
}

func StampCard() {
	usecase := user.InitializeUserCheckinUsecase()
	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	card, err := usecase.GetStampCard(userID)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(card.Checkins[0])
}
