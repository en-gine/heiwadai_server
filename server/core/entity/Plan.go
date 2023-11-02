package entity

import "github.com/google/uuid"

type Plan struct {
	ID        string
	Title     string
	Price     uint
	ImageURL  string
	RoomType  RoomType
	MealType  MealType
	SmokeType SmokeType
	OverView  string
	StoreID   uuid.UUID
}

type MealType struct {
	Morning bool
	Dinner  bool
}

func (m MealType) String() string {
	var mealType string
	var morning string
	var dinner string

	if m.Morning {
		morning = "朝食あり"
	} else {
		morning = "朝食なし"
	}
	if m.Dinner {
		dinner = "夕食あり"
	} else {
		dinner = "夕食なし"
	}
	mealType = morning + dinner
	if !m.Morning && !m.Dinner {
		mealType = "食事なし"
	}
	return mealType
}

type RoomType int

const (
	RoomTypeSingle RoomType = iota
	RoomTypeSemiDouble
	RoomTypeDouble
	RoomTypeTwin
	RoomTypeFourth
	RoomTypeUnknown
)

var RoomTypeAll = []RoomType{
	RoomTypeSingle,
	RoomTypeSemiDouble,
	RoomTypeDouble,
	RoomTypeTwin,
	RoomTypeFourth,
}

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
	case RoomTypeUnknown:
		fallthrough
	default:
		return "Unknown"
	}
}

func (s RoomType) Code() string {
	switch s {
	case RoomTypeSingle:
		return "Single"
	case RoomTypeSemiDouble:
		return "SemiDouble"
	case RoomTypeDouble:
		return "Double"
	case RoomTypeTwin:
		return "Twin"
	case RoomTypeFourth:
		return "Fourth"
	case RoomTypeUnknown:
		fallthrough
	default:
		return "Unknown"
	}
}

func IncludeRoomType(roomType []RoomType, target RoomType) bool {
	for _, v := range roomType {
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
	SmokeTypeUnknown
)

func (s SmokeType) String() string {
	switch s {
	case SmokeTypeNonSmoking:
		return "禁煙"
	case SmokeTypeSmoking:
		return "喫煙"
	case SmokeTypeUnknown:
		fallthrough
	default:
		return "Unknown"
	}
}

func IncludeSmokeType(smokeTypeArr []SmokeType, target SmokeType) bool {
	for _, v := range smokeTypeArr {
		if v == target {
			return true
		}
	}
	return false
}

func RegenPlan(
	ID string,
	Title string,
	Price uint,
	ImageURL string,
	RoomType RoomType,
	MealType MealType,
	SmokeType SmokeType,
	OverView string,
	StoreID uuid.UUID,
) *Plan {
	return &Plan{
		ID:        ID,
		Title:     Title,
		Price:     Price,
		ImageURL:  ImageURL,
		RoomType:  RoomType,
		MealType:  MealType,
		SmokeType: SmokeType,
		OverView:  OverView,
		StoreID:   StoreID,
	}
}
