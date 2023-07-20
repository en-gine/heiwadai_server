package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IPostRepository = &PostRepository{}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository() *PostRepository {
	db := InitDB()

	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) Save(updatePost *entity.Post) error {
	post := models.Post{
		ID:         updatePost.ID.String(),
		Title:      updatePost.Title,
		Content:    updatePost.Content,
		PostStatus: int(updatePost.PostStatus),
		PostDate:   updatePost.PostDate,
		CreateAt:   updatePost.CreateAt,
	}
	err := post.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	return err
}

func (pr *PostRepository) Delete(postId uuid.UUID) error {
	deletePost, err := models.FindPost(context.Background(), pr.db, postId.String())
	if err != nil {
		return err
	}
	_, err = deletePost.Delete(context.Background(), pr.db)
	return err
}
