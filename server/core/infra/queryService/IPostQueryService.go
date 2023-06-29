package queryservice

import (
	"server/core/entity"
)

type IPostQueryService interface {
	GetActiveAll(limit *int, page *int, per_page *int) ([]*entity.Post, error)
	GetAll(limit *int, page *int, per_page *int) ([]*entity.Post, error) //statusは不問
}
