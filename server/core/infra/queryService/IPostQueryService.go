package queryservice

import (
	"server/core/entity"
)

type IPostQueryService interface {
	GetByID(id int) (*entity.Post, error)
	GetAll() ([]*entity.Post, error) //statusは不問
}
