package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type ICheckinRepository interface {
	Save(tx ITransaction, checkin *entity.Checkin) error
	BulkArchive(tx ITransaction, userID uuid.UUID) error
}
