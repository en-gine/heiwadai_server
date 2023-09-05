package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IAdminRepository interface {
	Insert(admin *entity.Admin) error
	Update(admin *entity.Admin) error
	Delete(id uuid.UUID) error
}
