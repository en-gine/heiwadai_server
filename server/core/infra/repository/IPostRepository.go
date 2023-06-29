package repository

import (
	"server/core/entity"
)

type IPostRepository interface {
	Save(updatePost *entity.Post, user *entity.User) error
}
