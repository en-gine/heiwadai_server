package entity

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID           uuid.UUID
	CheckInDate  time.Time
	CheckOutDate time.Time
	ChackInTime  time.Time
	Content      string
	Cost         int
	Payment      string
	StayCustomer *StayCustomer
	StayStore    *Store
	ReservedPlan *Plan
	ReservedUser *User
	Note         string
}

type StayCustomer struct {
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   *string
	BirthDate     time.Time
	ZipCode       *string
	Prefecture    string
	City          *string
	Address       *string
	Tel           *string
	Mail          string
}

func CreateReservation(
	CheckInDate time.Time,
	CheckOutDate time.Time,
	ChackInTime time.Time,
	Content string,
	Cost int,
	Payment string,
	StayCustomer *StayCustomer,
	StayStore *Store,
	ReservedPlan *Plan,
	ReservedUser *User,
	Note string,
) *Reservation {
	return &Reservation{
		ID:           uuid.New(),
		CheckInDate:  CheckInDate,
		CheckOutDate: CheckOutDate,
		ChackInTime:  ChackInTime,
		Content:      Content,
		Cost:         Cost,
		Payment:      Payment,
		StayCustomer: StayCustomer,
		StayStore:    StayStore,
		ReservedPlan: ReservedPlan,
		ReservedUser: ReservedUser,
		Note:         Note,
	}
}
