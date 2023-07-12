package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"
)

type ICheckinQueryService interface {
	GetActiveCheckin(user *entity.User) ([]*entity.Checkin, error)
	GetLastStoreCheckin(user *entity.User, store *entity.Store) (*entity.Checkin, error) //前回の特定のStoreへのチェックイン
	GetAllCheckin(user *entity.User, pager *types.PageQuery) ([]*entity.Checkin, error)  //archiveかどうか不問
}
