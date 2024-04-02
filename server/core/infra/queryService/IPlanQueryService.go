package queryservice

import (
	"time"

	"server/core/entity"
)

type IPlanQueryService interface {
	Search(
		stayStore []*entity.StayableStore,
		stayFrom time.Time,
		stayTo time.Time,
		adult int,
		child int,
		roomCount int,
		smokeTypes *[]entity.SmokeType,
		mealType *entity.MealType,
		roomTypes *[]entity.RoomType,
	) (*[]entity.Plan, error)
}
