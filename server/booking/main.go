package main

import (
	"fmt"
	"time"

	"server/core/entity"
	"server/infrastructure/booking"
	implements "server/infrastructure/repository"
)

func main() {
	storeQuery := implements.NewStoreQueryService()
	p := booking.NewPlanQuery(storeQuery)

	tomorrow := time.Now().Add(2 * 24 * time.Hour)
	single := entity.RoomTypeSingle
	rooms := []entity.RoomType{single}
	meal := entity.MealType{Morning: true, Dinner: true}
	// smork := entity.SmokeTypeSmoking
	// smokes := []entity.SmokeType{smork}
	plans, err := p.Search(
		nil,
		time.Now(),
		tomorrow,
		1,
		1,
		1,
		nil,
		&meal,
		&rooms,
	)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(plans)
}
