package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IStoreQueryService interface {
	GetByID(id uuid.UUID) (*entity.Store, error)
	GetActiveAll() ([]*entity.Store, error)
	GetAll() ([]*entity.Store, error) //activeかどうか不問
}
