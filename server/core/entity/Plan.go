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
	Single RoomType = iota
	SemiDouble
	Double
	Twin
	Forth
)

func (s RoomType) String() string {
	switch s {
	case Single:
		return "シングル"
	case SemiDouble:
		return "セミダブル"
	case Double:
		return "ダブル"
	case Twin:
		return "ツイン"
	case Forth:
		return "フォース"
	default:
		return "Unknown"
	}
}

type SmokeType int

const (
	NonSmoking SmokeType = iota
	Smoking
)

func (s SmokeType) String() string {
	switch s {
	case NonSmoking:
		return "禁煙"
	case Smoking:
		return "喫煙"
	default:
		return "Unknown"
	}
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
		SmokeType: Smoking,
		OverView:  OverView,
	}
}
