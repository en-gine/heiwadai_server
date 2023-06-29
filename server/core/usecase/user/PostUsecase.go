package user

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
)

type PostUsecase struct {
	postRepository repository.IPostRepository
	postQuery      queryservice.IPostQueryService
}

func NewPostUsecase(postRepository repository.IPostRepository, postQuery queryservice.IPostQueryService) *PostUsecase {
	return &PostUsecase{
		postRepository: postRepository,
		postQuery:      postQuery,
	}
}

func (u *PostUsecase) GetActiveList() ([]*entity.Post, error) {

	posts, err := u.postQuery.GetActiveAll(nil, nil, nil)

	return posts, err
}
