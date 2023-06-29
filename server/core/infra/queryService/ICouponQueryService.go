package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type ICouponQueryService interface {
	GetById(id uuid.UUID) (*entity.Coupon, error)
	GetActiveAll(user *entity.User) ([]*entity.Coupon, error)
	GetAll(user *entity.User, limit *int, page *int, per_page *int) ([]*entity.Coupon, error) //使用済かどうか不問
}
