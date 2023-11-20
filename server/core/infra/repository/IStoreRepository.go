package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IStoreRepository interface {
	Save(store *entity.Store, stayableInfo *entity.StayableStoreInfo) error
	RegenQR(storeID uuid.UUID) (*uuid.UUID, error)
	RegenUnlimitQR(storeID uuid.UUID) (*uuid.UUID, error)
}
