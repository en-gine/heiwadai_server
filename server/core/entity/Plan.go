package entity

import (
	"github.com/google/uuid"
)

type Plan struct {
	ID        uuid.UUID
	Title     string
	Price     uint
	RoomType  RoomType
	MealType  MealType
	SmokeType SmokeType
	OverView  string
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

type MealType int

const (
	Morning MealType = iota
	Dinner
)

func (s MealType) String() string {
	switch s {
	case Morning:
		return "朝食あり"
	case Dinner:
		return "夕食あり"
	default:
		return "Unknown"
	}
}

func RegenPlan(
	RoomType RoomType,
	MealType MealType,
	SmokeType SmokeType,
	OverView string,

) *Plan {
	return &Plan{
		ID:        uuid.New(),
		RoomType:  RoomType,
		SmokeType: Smoking,
		OverView:  OverView,
	}
}
