package entity

import "time"

type PlanCalendar struct {
	Plan
	Date            time.Time
	Price           uint
	AvailableStatus AvailableStatus
}
type AvailableStatus int

const (
	Available AvailableStatus = iota
	NotAvailable
)

func (a AvailableStatus) String() string {
	switch a {
	case Available:
		return "Available"
	case NotAvailable:
		return "NotAvailable"
	default:
		return "Unknown"
	}
}
func NewPlanCalendar() *PlanCalendar {
	return &PlanCalendar{}
}
