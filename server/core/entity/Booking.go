package entity

import (
	"regexp"
	"time"

	"server/core/errors"

	"github.com/google/uuid"
)

type Booking struct {
	ID              uuid.UUID
	StayFrom        time.Time
	StayTo          time.Time
	Adult           uint
	Child           uint
	RoomCount       uint
	CheckInTime     CheckInTime
	TotalCost       uint
	GuestData       *GuestData
	BookPlan        *Plan
	BookUserID      uuid.UUID
	Note            string
	TlDataID        string  // TLDataID(予約番号)
	TlBookingNumber *string // TLBooking番号
}

type GuestData struct {
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   *string
	ZipCode       *string
	Prefecture    *Prefecture
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
	ZipCode *string,
	Prefecture *Prefecture,
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
	Prefecture *Prefecture,
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
	TlDataID string,
	TlBookingNumber *string,
) (*Booking, *errors.DomainError) {
	return &Booking{
		ID:              uuid.New(),
		StayFrom:        stayFrom,
		StayTo:          stayTo,
		Adult:           adult,
		Child:           child,
		RoomCount:       roomCount,
		CheckInTime:     CheckInTime,
		TotalCost:       TotalCost,
		GuestData:       GuestData,
		BookPlan:        BookPlan,
		BookUserID:      BookUserID,
		Note:            Note,
		TlDataID:        TlDataID,
		TlBookingNumber: TlBookingNumber,
	}, nil
}

type MaintenanceInfo struct {
	IsMaintenance bool
	Message       string
}

func GetBookingUnderMaintenance() *MaintenanceInfo {
	// 日本時間の毎月第2日曜26:00～28:30にTLリンカーンが定期メンテナンスを行う
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)

	year, month := now.Year(), now.Month()
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	weekday := firstDayOfMonth.Weekday()
	var secondSunday time.Time

	if weekday == time.Sunday {
		secondSunday = firstDayOfMonth.AddDate(0, 0, 7)
	} else {
		daysUntilFirstSunday := (7 - int(weekday)) % 7
		firstSunday := firstDayOfMonth.AddDate(0, 0, daysUntilFirstSunday)
		secondSunday = firstSunday.AddDate(0, 0, 7)
	}

	startMaintenance := time.Date(year, month, secondSunday.Day(), 2, 0, 0, 0, loc)
	endMaintenance := startMaintenance.Add(2*time.Hour + 30*time.Minute)

	if startMaintenance == endMaintenance {
		return &MaintenanceInfo{
			IsMaintenance: false,
			Message:       "",
		}
	}

	isMaintenace := now.After(startMaintenance) && now.Before(endMaintenance)
	message := "毎月第2日曜26:00～28:30は、\n予約システムの定期メンテナンスのため、\n予約及びキャンセル機能をご利用いただけません。"
	if isMaintenace {
		return &MaintenanceInfo{
			IsMaintenance: true,
			Message:       message,
		}
	}
	return &MaintenanceInfo{
		IsMaintenance: false,
		Message:       "",
	}
}
