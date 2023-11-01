package entity

import "time"

type PlanCalendar struct {
	Plan
	DateStatus []DateStatus
}

type DateStatus struct {
	Date            time.Time
	Price           uint
	AvailableStatus AvailableStatus
}
type AvailableStatus int

const (
	Available AvailableStatus = iota
	NotMuchLeft
	NotAvailable
)

func (a AvailableStatus) String() string {
	switch a {
	case Available:
		return "Available"
	case NotMuchLeft:
		return "NotMuchLeft"
	case NotAvailable:
		return "NotAvailable"
	default:
		return "Unknown"
	}
}

func NewPlanCalendar() *PlanCalendar {
	return &PlanCalendar{}
}
