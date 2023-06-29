package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Delete(id uuid.UUID) error
}
