package implement

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IPostQueryService = &PostQueryService{}

type PostQueryService struct {
	db *sql.DB
}

func NewPostQueryService() (*PostQueryService, error) {
	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	return &PostQueryService{
		db: db,
	}, nil
}

func (pq *PostQueryService) GetById(id uuid.UUID) (*entity.Post, error) {
	post, err := models.FindPost(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	postId := uuid.MustParse(post.ID)

	return &entity.Post{
		ID:         postId,
		Title:      post.Title,
		Content:    post.Content,
		PostStatus: entity.PostStatus(post.PostStatus),
		PostDate:   post.PostDate,
		CreateAt:   post.CreateAt,
	}, nil
}

func (pq *PostQueryService) GetActiveAll(pager *types.PageQuery) ([]*entity.Post, error) {
	posts, err := models.Posts(models.PostWhere.PostStatus.EQ(int(entity.PostPublish)), qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Post
	for _, post := range posts {
		postId := uuid.MustParse(post.ID)
		result = append(result, &entity.Post{
			ID:         postId,
			Title:      post.Title,
			Content:    post.Content,
			PostStatus: entity.PostStatus(post.PostStatus),
			PostDate:   post.PostDate,
			CreateAt:   post.CreateAt,
		})
	}
	return result, nil
}

func (pq *PostQueryService) GetAll(pager *types.PageQuery) ([]*entity.Post, error) {
	posts, err := models.Posts(qm.Load(models.PostRels.AuthorAdmin), qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Post
	for _, post := range posts {
		postId := uuid.MustParse(post.ID)
		result = append(result, &entity.Post{
			ID:         postId,
			Title:      post.Title,
			Content:    post.Content,
			PostStatus: entity.PostStatus(post.PostStatus),
			PostDate:   post.PostDate,
			CreateAt:   post.CreateAt,
		})
	}
	return result, nil

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
