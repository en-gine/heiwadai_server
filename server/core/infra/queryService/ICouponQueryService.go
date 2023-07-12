package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type ICouponQueryService interface {
	GetById(id uuid.UUID) (*entity.Coupon, error)
	GetActiveAll(user *entity.User) ([]*entity.Coupon, error)
	GetAll(user *entity.User, pager *types.PageQuery) ([]*entity.Coupon, error) //使用済かどうか不問
}
