package entity

import (
	"regexp"
	"server/core/errors"
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID           uuid.UUID
	CheckInDate  time.Time
	CheckOutDate time.Time
	CheckInTime  CheckInTime
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

type CheckInTime string

func NewCheckInTime(s string) (*CheckInTime, *errors.DomainError) {
	if IsValidTimeFormat(s) {
		return nil, errors.NewDomainError(errors.InvalidParameter, "CheckInTimeの形式が正しくありません。")
	}
	result := CheckInTime(s)
	return &result, nil
}

func IsValidTimeFormat(s string) bool {
	// パターンは 00~23の時と 00~59の分にマッチします
	pattern := `^([01]?[0-9]|2[0-3]):[0-5][0-9]$`
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func CreateReservation(
	CheckInDate time.Time,
	CheckOutDate time.Time,
	CheckInTime CheckInTime,
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
		CheckInTime:  CheckInTime,
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
