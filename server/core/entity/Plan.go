package entity

import (
	"github.com/google/uuid"
)

type Plan struct {
	ID        uuid.UUID
	Title     string
	Price     uint
	ImageURL  string
	RoomType  RoomType
	MealType  MealType
	SmokeType SmokeType
	OverView  string
}

type MealType struct {
	Morning bool
	Dinner  bool
}

func (m MealType) String() string {
	var mealType string

	morning := "朝食なし"
	dinner := "夕食なし"
	if m.Morning {
		morning = "朝食あり"
	}
	if m.Dinner {
		dinner = "夕食あり"
	}
	mealType = morning + dinner
	return mealType
}

type RoomType int

const (
	RoomTypeSingle RoomType = iota
	RoomTypeSemiDouble
	RoomTypeDouble
	RoomTypeTwin
	RoomTypeFourth
)

func (s RoomType) String() string {
	switch s {
	case RoomTypeSingle:
		return "シングル"
	case RoomTypeSemiDouble:
		return "セミダブル"
	case RoomTypeDouble:
		return "ダブル"
	case RoomTypeTwin:
		return "ツイン"
	case RoomTypeFourth:
		return "フォース"
	default:
		return "Unknown"
	}
}

func IncludeRoomType(roomType *[]RoomType, target RoomType) bool {
	for _, v := range *roomType {
		if v == target {
			return true
		}
	}
	return false
}

type SmokeType int

const (
	SmokeTypeNonSmoking SmokeType = iota
	SmokeTypeSmoking
)

func (s SmokeType) String() string {
	switch s {
	case SmokeTypeNonSmoking:
		return "禁煙"
	case SmokeTypeSmoking:
		return "喫煙"
	default:
		return "Unknown"
	}
}

func IncludeSmokeType(smokeType *[]SmokeType, target SmokeType) bool {
	for _, v := range *smokeType {
		if v == target {
			return true
		}
	}
	return false
}

func RegenPlan(
	RoomType RoomType,
	MealType MealType,
	ImageURL string,
	SmokeType SmokeType,
	OverView string,

) *Plan {
	return &Plan{
		ID:        uuid.New(),
		RoomType:  RoomType,
		MealType:  MealType,
		ImageURL:  ImageURL,
		SmokeType: SmokeType,
		OverView:  OverView,
	}
}
