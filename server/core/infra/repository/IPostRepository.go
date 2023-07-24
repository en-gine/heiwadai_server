package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IPostRepository interface {
	Save(updatePost *entity.Post) error
	Delete(postID uuid.UUID) error
}
