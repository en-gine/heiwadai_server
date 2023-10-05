package repository

import (
	"context"

	"server/core/entity"
)

type IUserCouponRepository interface {
	Save(ctx context.Context, coupon *entity.UserAttachedCoupon) error
}
