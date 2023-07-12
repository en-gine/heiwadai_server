package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
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

func (u *PostUsecase) GetList(pager *types.PageQuery) ([]*entity.Post, *errors.DomainError) {

	posts, err := u.postQuery.GetAll(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return posts, nil
}

func (u *PostUsecase) Create(title string, content string, postDate time.Time, auther entity.Admin) (*entity.Post, *errors.DomainError) {

	post := entity.CreatePost(title, content, postDate, auther)

	err := u.postRepository.Save(post)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return post, nil
}

func (u *PostUsecase) Update(title *string, content *string, postDate *time.Time, auther entity.Admin, postId uuid.UUID) (*entity.Post, *errors.DomainError) {

	oldPost, err := u.postQuery.GetById(postId)

	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if oldPost == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}
	var updateTitle string
	var updateContent string
	var updatePostDate time.Time
	if title != nil {
		updateTitle = *title
	} else {
		updateTitle = oldPost.Title
	}

	if content != nil {
		updateContent = *content
	} else {
		updateContent = oldPost.Content
	}

	if postDate != nil {
		updatePostDate = *postDate
	} else {
		updatePostDate = oldPost.PostDate
	}

	updatePost := entity.UpdatePost(postId, updateTitle, updateContent, updatePostDate, auther)

	err = u.postRepository.Save(updatePost)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updatePost, nil
}

func (u *PostUsecase) Delete(postId uuid.UUID) (*entity.Post, *errors.DomainError) {
	deletePost, err := u.postQuery.GetById(postId)

	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if deletePost == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}

	err = u.postRepository.Delete(postId)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return deletePost, nil
}
