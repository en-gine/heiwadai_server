package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMealType_String(t *testing.T) {
	tests := []struct {
		name     string
		mealType MealType
		want     string
	}{
		{"両方あり", MealType{Morning: true, Dinner: true}, "朝食あり夕食あり"},
		{"朝食のみ", MealType{Morning: true, Dinner: false}, "朝食あり夕食なし"},
		{"夕食のみ", MealType{Morning: false, Dinner: true}, "朝食なし夕食あり"},
		{"両方なし", MealType{Morning: false, Dinner: false}, "食事なし"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.mealType.String())
		})
	}
}

func TestRoomType_String(t *testing.T) {
	tests := []struct {
		roomType RoomType
		want     string
	}{
		{RoomTypeSingle, "シングル"},
		{RoomTypeSemiDouble, "セミダブル"},
		{RoomTypeDouble, "ダブル"},
		{RoomTypeTwin, "ツイン"},
		{RoomTypeFourth, "フォース"},
		{RoomTypeUnknown, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.roomType.String())
		})
	}
}

func TestRoomType_Code(t *testing.T) {
	tests := []struct {
		roomType RoomType
		want     string
	}{
		{RoomTypeSingle, "Single"},
		{RoomTypeSemiDouble, "SemiDouble"},
		{RoomTypeDouble, "Double"},
		{RoomTypeTwin, "Twin"},
		{RoomTypeFourth, "Fourth"},
		{RoomTypeUnknown, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.roomType.Code())
		})
	}
}

func TestIncludeRoomType(t *testing.T) {
	roomTypes := []RoomType{RoomTypeSingle, RoomTypeDouble}

	assert.True(t, IncludeRoomType(roomTypes, RoomTypeSingle))
	assert.True(t, IncludeRoomType(roomTypes, RoomTypeDouble))
	assert.False(t, IncludeRoomType(roomTypes, RoomTypeTwin))
}

func TestSmokeType_String(t *testing.T) {
	tests := []struct {
		smokeType SmokeType
		want      string
	}{
		{SmokeTypeNonSmoking, "禁煙"},
		{SmokeTypeSmoking, "喫煙"},
		{SmokeTypeUnknown, "情報なし"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.smokeType.String())
		})
	}
}

func TestIncludeSmokeType(t *testing.T) {
	smokeTypes := []SmokeType{SmokeTypeNonSmoking, SmokeTypeSmoking}

	assert.True(t, IncludeSmokeType(smokeTypes, SmokeTypeNonSmoking))
	assert.True(t, IncludeSmokeType(smokeTypes, SmokeTypeSmoking))
	assert.False(t, IncludeSmokeType(smokeTypes, SmokeTypeUnknown))
}

func TestNewPlanCandidate(t *testing.T) {
	plan := &Plan{
		ID:    "1",
		Title: "Test Plan",
		Price: 10000,
	}

	planCandidate := NewPlanCandidate(plan, 2, 2)

	assert.Equal(t, plan, planCandidate.Plan)
	assert.Equal(t, uint(2500), planCandidate.MinimumPrice)
	assert.Equal(t, PricePerNightAndPerson, planCandidate.PricePerCategory)
}

func TestRegenPlan(t *testing.T) {
	storeID := uuid.New()
	plan := RegenPlan(
		"1",
		"Test Plan",
		10000,
		"http://example.com/image.jpg",
		RoomTypeSingle,
		MealType{Morning: true, Dinner: false},
		SmokeTypeNonSmoking,
		"Test Overview",
		storeID,
		"TEST_ROOM_CODE",
		"TEST_ROOM_NAME",
	)

	assert.Equal(t, "1", plan.ID)
	assert.Equal(t, "Test Plan", plan.Title)
	assert.Equal(t, uint(10000), plan.Price)
	assert.Equal(t, "http://example.com/image.jpg", plan.ImageURL)
	assert.Equal(t, RoomTypeSingle, plan.RoomType)
	assert.Equal(t, MealType{Morning: true, Dinner: false}, plan.MealType)
	assert.Equal(t, SmokeTypeNonSmoking, plan.SmokeType)
	assert.Equal(t, "Test Overview", plan.OverView)
	assert.Equal(t, storeID, plan.StoreID)
	assert.Equal(t, "TEST_ROOM_CODE", plan.TlBookingRoomTypeCode)
	assert.Equal(t, "TEST_ROOM_NAME", plan.TlBookingRoomTypeName)
}
