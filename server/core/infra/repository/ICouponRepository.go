package repository

import (
	"server/core/entity"
)

type ICouponRepository interface {
	Save(tx ITransaction, coupon *entity.Coupon) error
}
