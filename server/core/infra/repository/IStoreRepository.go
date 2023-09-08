package repository

import (
	"server/core/entity"
)

type IStoreRepository interface {
	Save(store *entity.Store, stayableInfo *entity.StayableStoreInfo) error
}
