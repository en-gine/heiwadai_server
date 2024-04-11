package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Save(user *entity.User, userOption *entity.UserOption) error
	Delete(id uuid.UUID) error
	DeleteUnderRegisterUser(id uuid.UUID) error
}
