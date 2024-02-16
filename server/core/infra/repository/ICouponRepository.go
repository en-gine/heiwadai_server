package repository

import (
	"context"

	"server/core/entity"
)

type ICouponRepository interface {
	Save(ctx context.Context, coupon *entity.Coupon) error
}
