package queryservice

import (
	"server/core/entity"
)

type ICheckinQueryService interface {
	GetActiveCheckin(user *entity.User) ([]*entity.Checkin, error)
	GetLastStoreCheckin(user *entity.User, store entity.Store) (*entity.Checkin, error)
	GetAllCheckin(user *entity.User, limit *int, page *int, per_page *int) ([]*entity.Checkin, error) //archiveかどうか不問
}
