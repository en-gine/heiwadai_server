package queryservice

import (
	"server/core/entity"
	"time"
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
	) ([]*entity.Plan, error)
}
