package main

import (
	"fmt"

	"server/infrastructure/logger"
	"server/router/user"

	"github.com/google/uuid"
)

func main() {
	// StampCard()
	Coupon()
	// Checkin()
}

func Coupon() {
	usecase := user.InitializeUserCouponUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	couponID, _ := uuid.Parse("5b3348d1-4f06-46f1-9fc5-75ca35285654")
	coupons, err := usecase.GetMyList(userID)
	if err != nil {
		logger.Error(err.Error())
	}

	for _, coupon := range coupons {
		fmt.Printf("%+v\n", coupon)
	}

	coupon, err := usecase.GetByID(userID, couponID)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Printf("%+v\n", coupon)

	// err = usecase.UseMyCoupon(userID, couponID)
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
}
