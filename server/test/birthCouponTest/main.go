package main

import (
	"fmt"
	"os"
	"time"

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

func ExpireAt() {
	os.Setenv("TZ", "Asia/Tokyo")
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	firstDayOfNextMonth := time.Date(currentYear, currentMonth+1, 1, 0, 0, 0, 0, now.Location())
	lastDayOfMonth := firstDayOfNextMonth.Add(-24 * time.Hour)
	fmt.Println(lastDayOfMonth)

	// ExpireAt := time.Now().AddDate(0, 1, -time.Now().Day()).AddDate(0, 0, -1).Truncate(24 * time.Hour)
	// fmt.Println(ExpireAt)
}
