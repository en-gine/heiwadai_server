package queryservice

import (
	"time"

	"server/core/entity"

	"github.com/google/uuid"
)

type IPlanQueryService interface {
	Search(
		stayStore []entity.Store,
		stayFrom time.Time,
		stayTo time.Time,
		adult int,
		child int,
		roomCount int,
		smokeTypes *[]entity.SmokeType,
		mealType *entity.MealType,
		roomTypes *[]entity.RoomType,
	) (*[]entity.Plan, error)
	GetMyBooking(userID uuid.UUID) (*[]entity.Plan, error)
}
