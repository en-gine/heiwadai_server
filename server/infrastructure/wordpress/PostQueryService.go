package repository

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
)

var _ queryservice.IPostQueryService = &PostQueryService{}

type PostQueryService struct {
}

func NewPostQueryService() *PostQueryService {

	return &PostQueryService{}
}

func (pq *PostQueryService) GetByID(id uuid.UUID) (*entity.Post, error) {

	return &entity.Post{}, nil
}

func (pq *PostQueryService) GetActiveAll(pager *types.PageQuery) ([]*entity.Post, error) {

	return nil, nil
}

func (pq *PostQueryService) GetAll(pager *types.PageQuery) ([]*entity.Post, error) {
	return nil, nil

}

func PostModelToEntity(post *models.Post, author entity.Admin) *entity.Post {
	return entity.RegenPost(
		uuid.MustParse(post.ID),
		post.Title,
		post.Content,
		author,
		entity.PostStatus(post.PostStatus),
		post.PostDate,
		post.CreateAt,
		post.UpdateAt,
	)
}
