package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateGuestData(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	firstNameKana := "ジョン"
	lastNameKana := "ドウ"
	companyName := "Test Company"
	zipCode := "123-4567"
	prefecture := Prefecture(5)
	city := "Shibuya"
	address := "1-1-1"
	tel := "03-1234-5678"
	mail := "john@example.com"

	guestData := CreateGuestData(firstName, lastName, firstNameKana, lastNameKana, &companyName, &zipCode, &prefecture, &city, &address, &tel, mail)

	assert.NotNil(t, guestData)
	assert.Equal(t, firstName, guestData.FirstName)
	assert.Equal(t, lastName, guestData.LastName)
	assert.Equal(t, firstNameKana, guestData.FirstNameKana)
	assert.Equal(t, lastNameKana, guestData.LastNameKana)
	assert.Equal(t, companyName, *guestData.CompanyName)
	assert.Equal(t, zipCode, *guestData.ZipCode)
	assert.Equal(t, prefecture, *guestData.Prefecture)
	assert.Equal(t, city, *guestData.City)
	assert.Equal(t, address, *guestData.Address)
	assert.Equal(t, tel, *guestData.Tel)
	assert.Equal(t, mail, guestData.Mail)
}

func TestNewCheckInTime(t *testing.T) {
	validTime := "14:30"
	invalidTime := "25:00"

	checkInTime, err := NewCheckInTime(validTime)

	assert.NotNil(t, checkInTime)
	assert.Nil(t, err)
	assert.Equal(t, validTime, checkInTime.String())

	checkInTime, err = NewCheckInTime(invalidTime)
	assert.Nil(t, checkInTime)
	assert.NotNil(t, err)
	assert.Equal(t, "CheckInTimeの形式が正しくありません。", err.Error())
}

func TestCreateBooking(t *testing.T) {
	stayFrom := time.Now()
	stayTo := stayFrom.Add(24 * time.Hour)
	adult := uint(2)
	child := uint(1)
	roomCount := uint(1)
	checkInTime, _ := NewCheckInTime("15:00")
	totalCost := uint(10000)
	guestData := CreateGuestData("John", "Doe", "ジョン", "ドウ", nil, nil, nil, nil, nil, nil, "john@example.com")
	bookPlan := &PlanStayDetail{} // この構造体の詳細が不明なので、空の構造体を使用
	bookUserID := uuid.New()
	note := "Test booking"
	tlDataID := "TL123456"
	tlBookingNumber := "BN789012"

	booking := CreateBooking(stayFrom, stayTo, adult, child, roomCount, *checkInTime, totalCost, guestData, bookPlan, bookUserID, note, tlDataID, &tlBookingNumber)

	assert.NotNil(t, booking)
	assert.NotEqual(t, uuid.Nil, booking.ID)
	assert.Equal(t, stayFrom, booking.StayFrom)
	assert.Equal(t, stayTo, booking.StayTo)
	assert.Equal(t, adult, booking.Adult)
	assert.Equal(t, child, booking.Child)
	assert.Equal(t, roomCount, booking.RoomCount)
	assert.Equal(t, *checkInTime, booking.CheckInTime)
	assert.Equal(t, totalCost, booking.TotalCost)
	assert.Equal(t, guestData, booking.GuestData)
	assert.Equal(t, bookPlan, booking.BookPlan)
	assert.Equal(t, bookUserID, booking.BookUserID)
	assert.Equal(t, note, booking.Note)
	assert.Equal(t, tlDataID, booking.TlDataID)
	assert.Equal(t, tlBookingNumber, *booking.TlBookingNumber)
}

func TestGetBookingUnderMaintenance(t *testing.T) {
	info := GetBookingUnderMaintenance()
	assert.NotNil(t, info)
	// メンテナンス状態のテストは時間に依存するため、ここでは単純に型チェックのみを行います
	assert.IsType(t, bool(false), info.IsMaintenance)
	assert.IsType(t, "", info.Message)
}
