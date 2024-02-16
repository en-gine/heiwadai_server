package repository

import (
	"context"

	"server/core/entity"
)

type IUserCouponRepository interface {
	Save(ctx context.Context, coupon *entity.UserAttachedCoupon) error
	IssueAll(ctx context.Context, coupon *entity.Coupon) (int, error) // 発行数を返す
}
