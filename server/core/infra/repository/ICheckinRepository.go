package repository

import (
	"context"

	"server/core/entity"
)

type ICheckinRepository interface {
	Save(ctx context.Context, checkin *entity.Checkin) error
}
