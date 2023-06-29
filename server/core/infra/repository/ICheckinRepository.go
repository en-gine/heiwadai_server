package repository

import (
	"server/core/entity"
)

type ICheckinRepository interface {
	Save(checkin *entity.Checkin) error
}
