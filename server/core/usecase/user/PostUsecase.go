package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
)

type PostUsecase struct {
	postQuery queryservice.IPostQueryService
}

func NewPostUsecase(postQuery queryservice.IPostQueryService) *PostUsecase {
	return &PostUsecase{
		postQuery: postQuery,
	}
}

func (u *PostUsecase) GetList() ([]*entity.Post, *errors.DomainError) {

	posts, err := u.postQuery.GetAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return posts, nil
}

func (u *PostUsecase) GetByID(postID uint32) (*entity.Post, *errors.DomainError) {

	post, err := u.postQuery.GetByID(int(postID))
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return post, nil
}
