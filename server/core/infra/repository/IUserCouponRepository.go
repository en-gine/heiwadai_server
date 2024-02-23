package repository

import (
	"server/core/entity"
)

type IUserCouponRepository interface {
	Save(tx ITransaction, coupon *entity.UserAttachedCoupon) error
	IssueAll(tx ITransaction, coupon *entity.Coupon, birthMonth *int) (int, error) // 発行数を返す
}
