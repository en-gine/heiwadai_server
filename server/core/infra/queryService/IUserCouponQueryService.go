package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserCouponQueryService interface {
	GetById(user *entity.User, couponid uuid.UUID) (*entity.UserAttachedCoupon, error)
	GetActiveAll(user *entity.User) ([]*entity.UserAttachedCoupon, error)
	GetAll(user *entity.User, pager *types.PageQuery) ([]*entity.UserAttachedCoupon, error) //使用済かどうか不問
}
