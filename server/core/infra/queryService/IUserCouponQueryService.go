package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserCouponQueryService interface {
	GetByID(userID uuid.UUID, couponID uuid.UUID) (*entity.UserAttachedCoupon, error)
	GetActiveAll(userID uuid.UUID) ([]*entity.UserAttachedCoupon, error)
	GetExpires(userID uuid.UUID, limit int) ([]*entity.UserAttachedCoupon, error)
	GetUseds(userID uuid.UUID, limit int) ([]*entity.UserAttachedCoupon, error)
	GetAll(userID uuid.UUID, pager *types.PageQuery) ([]*entity.UserAttachedCoupon, *types.PageResponse, error) // 使用済かどうか不問
}
