package main

import (
	adminRouter "server/router/admin"
)

func CouponSeeder() {
	couponUsecase := adminRouter.InitializeAdminCouponUsecase()

	err := couponUsecase.CreateStandardCoupon()
	if err != nil {
		panic(err)
	}
}
