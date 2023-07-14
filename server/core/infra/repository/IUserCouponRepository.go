package repository

import (
	"server/core/entity"
)

type IUserCouponRepository interface {
	Save(coupon *entity.UserAttachedCoupon) error
}
