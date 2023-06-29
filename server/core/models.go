package entity

import (
	"time"

	"github.com/google/uuid"
)

// アプリユーザー
type User struct {
	ID            uuid.UUID
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
	AcceptMail    bool
}

// アプリユーザーのオプション（バックエンドのみで使用）
type UserOption struct {
	ID              uuid.UUID
	InnerNote       string
	IsBlackCustomer bool
}
type UserWithOption struct {
	UserOption
	User
}

// 　管理画面ユーザー
type Admin struct {
	ID       uuid.UUID
	Name     string
	BelongTo Store
}

// チェックイン
type Checkin struct {
	ID uuid.UUID
	Store
	User
	CheckInAt time.Time
	Archive   bool //archive=trueはログ的に管理画面側で確認するためのものです。Flutter側では気にする必要はありません。
}

// クーポン
type Coupon struct {
	ID                uuid.UUID
	Name              string
	CouponType        CouponType // "Standard" "Custom" "Birthday"
	DiscountAmount    uint       //割引額
	ExpireAt          time.Time  //有効期限
	IsCombinationable bool       //併用可能
	OverView          string     //概要 今の
	UsedAt            *time.Time //使用済
	User              *User
	TargetStore       []*Store //対象店舗
}

type CouponType int

const (
	Standard CouponType = iota
	Custom
	Birthday
)

func (s CouponType) String() string {
	switch s {
	case Standard:
		return "Standard"
	case Custom:
		return "Custom"
	case Birthday:
		return "Birthday"
	default:
		return "Unknown"
	}
}

// 店舗

type Store struct {
	ID              uuid.UUID
	Name            string
	Address         string
	IsActive        bool
	StayAble        bool //宿泊施設かどうか
	QrCode          uuid.UUID
	UnLimitedQrCode uuid.UUID
}

// TOPバナー　WordPressから取ってくるかも
type Banner struct {
	ID       uuid.UUID
	ImageURL string
	Url      string
}

// お知らせ WordPressから取ってくるかも
type Post struct {
	ID         uuid.UUID
	Title      string
	Content    string
	Author     Admin
	PostStatus PostStatus //下書きor公開。flutter側では気にする必要なし
}
type PostStatus int

const (
	Draft PostStatus = iota
	Publish
)

// 予約
type Reservation struct {
	ID           uuid.UUID
	CheckIn      time.Time
	CheckOut     time.Time
	ChackInTime  time.Time
	Content      string
	Cost         int
	Payment      string
	StayCustomer *StayCustomer //宿泊者情報
	StayStore    *Store        //滞在予定店舗
	ReservedPlan *Plan         //予約したプラン
	ReservedUser *User         //予約したユーザー
	Note         string
}

// 宿泊プラン
type Plan struct {
	ID           uuid.UUID
	RoomType     string
	IncludedMeal string
	Smoking      string
	OverView     string
}
type PlanCalendar struct {
	Plan
	Date    time.Time
	Price   uint
	Vacancy string //空状況
}

// 宿泊者情報
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
