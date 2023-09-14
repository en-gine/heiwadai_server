package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type ICheckinQueryService interface {
	GetActiveCheckin(userID uuid.UUID) ([]*entity.Checkin, error)
	GetLastStoreCheckin(userID uuid.UUID, storeID uuid.UUID) (*entity.Checkin, error)  //前回の特定のStoreへのチェックイン
	GetAllCheckin(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, error) //archiveかどうか不問
}
