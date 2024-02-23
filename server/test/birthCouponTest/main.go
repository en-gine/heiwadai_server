package main

import (
	"fmt"

	"server/infrastructure/logger"
	"server/router/admin"
)

func main() {
	Coupon()
}

func Coupon() {
	usecase := admin.InitializeAdminCouponUsecase()

	count, err := usecase.BulkAttachBirthdayCoupon(1)
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Printf("%+v\n", count)
}
