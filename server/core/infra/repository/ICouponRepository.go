package repository

import (
	"server/core/entity"
)

type ICouponRepository interface {
	Save(coupon *entity.Coupon) error
}
