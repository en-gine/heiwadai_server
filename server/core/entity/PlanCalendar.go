package entity

import "time"

type PlanCalendar struct {
	Plan
	Date    time.Time
	Price   uint
	Vacancy bool
}

func NewPlanCalendar() *PlanCalendar {
	return &PlanCalendar{}
}
