package entity

import (
	"github.com/google/uuid"
)

type Plan struct {
	ID           uuid.UUID
	RoomType     string
	IncludedMeal string
	Smoking      string
	OverView     string
}

func RegenPlan(
	RoomType string,
	IncludedMeal string,
	Smoking string,
	OverView string,

) *Plan {
	return &Plan{
		ID:       uuid.New(),
		RoomType: RoomType,
		Smoking:  Smoking,
		OverView: OverView,
	}
}
