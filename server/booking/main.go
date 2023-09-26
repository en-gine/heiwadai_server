package main

import (
	"fmt"
	"server/infrastructure/booking"
	"time"
)

func main() {
	p := booking.NewPlanQuery()

	tomorrow := time.Now().Add(2 * 24 * time.Hour)
	plans, err := p.Search(
		nil,
		time.Now(),
		tomorrow,
		1,
		1,
		1,
		nil,
		nil,
		nil,
	)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(plans)
}
