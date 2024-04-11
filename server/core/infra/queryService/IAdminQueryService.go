package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IAdminQueryService interface {
	GetByID(id uuid.UUID) (*entity.Admin, error)
	GetByMail(mail entity.Mail) (*entity.Admin, error)
	GetAll() ([]*entity.Admin, error)
}
