package entity

import (
	"regexp"
	"time"

	"server/core/errors"

	"github.com/google/uuid"
)

type Booking struct {
	ID           uuid.UUID
	BookSystemID string // TLBooking番号
	StayFrom     time.Time
	StayTo       time.Time
	Adult        uint
	Child        uint
	RoomCount    uint
	CheckInTime  CheckInTime
	TotalCost    uint
	GuestData    *GuestData
	BookPlan     *Plan
	BookUserID   uuid.UUID
	Note         string
}

type GuestData struct {
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   *string
	BirthDate     time.Time
	ZipCode       *string
	Prefecture    Prefecture
	City          *string
	Address       *string
	Tel           *string
	Mail          string
}

func CreateGuestData(
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture Prefecture,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
) *GuestData {
	return &GuestData{
		FirstName:     FirstName,
		LastName:      LastName,
		FirstNameKana: FirstNameKana,
		LastNameKana:  LastNameKana,
		CompanyName:   CompanyName,
		BirthDate:     BirthDate,
		ZipCode:       ZipCode,
		Prefecture:    Prefecture,
		City:          City,
		Address:       Address,
		Tel:           Tel,
		Mail:          Mail,
	}
}

func RegenGuestData(
	ID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture Prefecture,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool,
) *GuestData {
	return &GuestData{
		FirstName:     FirstName,
		LastName:      LastName,
		FirstNameKana: FirstNameKana,
		LastNameKana:  LastNameKana,
		CompanyName:   CompanyName,
		BirthDate:     BirthDate,
		ZipCode:       ZipCode,
		Prefecture:    Prefecture,
		City:          City,
		Address:       Address,
		Tel:           Tel,
		Mail:          Mail,
	}
}

type CheckInTime string

func NewCheckInTime(s string) (*CheckInTime, *errors.DomainError) {
	if IsValidTimeFormat(s) {
		return nil, errors.NewDomainError(errors.InvalidParameter, "CheckInTimeの形式が正しくありません。")
	}
	result := CheckInTime(s)
	return &result, nil
}

func (c *CheckInTime) String() string {
	return string(*c)
}

func IsValidTimeFormat(s string) bool {
	// パターンは 00~23の時と 00~59の分にマッチします
	pattern := `^([01]?[0-9]|2[0-3]):[0-5][0-9]$`
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func CreateBooking(
	stayFrom time.Time,
	stayTo time.Time,
	adult uint,
	child uint,
	roomCount uint,
	CheckInTime CheckInTime,
	TotalCost uint,
	GuestData *GuestData,
	BookPlan *Plan,
	BookUserID uuid.UUID,
	Note string,
) *Booking {
	return &Booking{
		ID:          uuid.New(),
		StayFrom:    stayFrom,
		StayTo:      stayTo,
		Adult:       adult,
		Child:       child,
		RoomCount:   roomCount,
		CheckInTime: CheckInTime,
		TotalCost:   TotalCost,
		GuestData:   GuestData,
		BookPlan:    BookPlan,
		BookUserID:  BookUserID,
		Note:        Note,
	}
}
