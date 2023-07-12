package user

import (
	"server/core/entity"
	"server/core/errors"
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

func (u *PostUsecase) GetActiveList() ([]*entity.Post, *errors.DomainError) {

	posts, err := u.postQuery.GetActiveAll(nil)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return posts, nil
}
