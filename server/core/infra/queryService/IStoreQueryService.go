package queryservice

import (
	"server/core/entity"
)

type IStoreQueryService interface {
	GetActiveAll() ([]*entity.Store, error)
	GetAll(limit *int, page *int, per_page *int) ([]*entity.Store, error) //activeかどうか不問
}
