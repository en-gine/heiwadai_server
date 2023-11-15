package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type ICheckinQueryService interface {
	GetMyActiveCheckin(userID uuid.UUID) ([]*entity.Checkin, error)
	GetMyLastStoreCheckin(userID uuid.UUID, storeID uuid.UUID) (*entity.Checkin, error)  // 前回の特定のStoreへのチェックイン
	GetMyAllCheckin(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, error) // archiveかどうか不問
	GetAllUserAllCheckin(pager *types.PageQuery) ([]*entity.Checkin, error)              // 前回のユーザーチェックインarchiveかどうか不問
}
