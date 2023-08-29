package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IMailMagazineRepository interface {
	Save(updateMailMagazine *entity.MailMagazine) error
	Delete(magazineID uuid.UUID) error
}
