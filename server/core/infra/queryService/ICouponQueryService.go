package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type ICouponQueryService interface {
	GetById(id uuid.UUID) (*entity.Coupon, error)
	GetCouponByType(couponType entity.CouponType) (*entity.Coupon, error)
	GetCouponListByType(couponType entity.CouponType, pager *types.PageQuery) ([]*entity.Coupon, error)
}
