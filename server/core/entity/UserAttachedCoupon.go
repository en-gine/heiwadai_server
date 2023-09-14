package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserAttachedCoupon struct {
	UserID uuid.UUID
	*Coupon
	UsedAt *time.Time
}

func CreateUserAttachedCoupon(userID uuid.UUID, coupon *Coupon) *UserAttachedCoupon {
	return &UserAttachedCoupon{
		UserID: userID,
		Coupon: coupon,
		UsedAt: nil,
	}
}

func UseUserAttachedCoupon(userID uuid.UUID, coupon *Coupon) *UserAttachedCoupon {
	now := time.Now()
	return &UserAttachedCoupon{
		UserID: userID,
		Coupon: coupon,
		UsedAt: &now,
	}
}

func RegenUserAttachedCoupon(userID uuid.UUID, coupon *Coupon, useAt *time.Time) *UserAttachedCoupon {

	return &UserAttachedCoupon{
		UserID: userID,
		Coupon: coupon,
		UsedAt: useAt,
	}
}
