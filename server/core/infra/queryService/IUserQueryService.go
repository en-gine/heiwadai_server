package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IUserQueryService interface {
	GetById(id uuid.UUID) (*entity.User, error)
	GetByMail(mail string) (*entity.User, error)
}
