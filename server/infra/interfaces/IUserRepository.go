package interfaces

import (
	"server/domain"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	FindById(uuid uuid.UUID) (*domain.User, error)
	FindByMail(mail string) (*domain.User, error)
}
