package entity

import "time"

type UserAttachedCoupon struct {
	*User
	*Coupon
	UsedAt *time.Time
}

func CreateUserAttachedCoupon(user *User, coupon *Coupon) *UserAttachedCoupon {
	return &UserAttachedCoupon{
		User:   user,
		Coupon: coupon,
		UsedAt: nil,
	}
}

func UseUserAttachedCoupon(user *User, coupon *Coupon) *UserAttachedCoupon {
	now := time.Now()
	return &UserAttachedCoupon{
		User:   user,
		Coupon: coupon,
		UsedAt: &now,
	}
}

func RegenUserAttachedCoupon(user *User, coupon *Coupon, useAt *time.Time) *UserAttachedCoupon {

	return &UserAttachedCoupon{
		User:   user,
		Coupon: coupon,
		UsedAt: useAt,
	}
}
