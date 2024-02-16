package repository

import (
	"context"

	"server/core/entity"

	"github.com/google/uuid"
)

type ICheckinRepository interface {
	Save(ctx context.Context, checkin *entity.Checkin) error
	BulkArchive(ctx context.Context, userID uuid.UUID) error
}
