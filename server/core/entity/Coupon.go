package entity

import (
	"server/core/errors"
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID                uuid.UUID
	Name              string
	CouponType        CouponType
	DiscountAmount    uint       //割引額
	ExpireAt          time.Time  //有効期限
	IsCombinationable bool       //併用可能
	Notices           []string   //注意事項
	UsedAt            *time.Time //使用済
	User              *User
	TargetStore       []*Store //対象店舗
	CreateAt          time.Time
	Status            CouponStatus
}

var DefaultNotices = []string{"クーポンは併用できません", "ランチではお使いになれません"}

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

type CouponStatus int

const (
	CouponDraft CouponStatus = iota
	CouponSaved
	CouponIssued
	CouponUsed
)

func (b CouponStatus) String() string {
	switch b {
	case CouponDraft:
		return "Draft"
	case CouponSaved:
		return "Saved"
	case CouponIssued:
		return "Issued"
	case CouponUsed:
		return "Used"
	default:
		return "Unknown"
	}
}

func newCoupon(
	ID uuid.UUID,
	Name string,
	CouponType CouponType,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	User *User,
	TargetStore []*Store,
	CreateAt time.Time,
	Status CouponStatus,
) (*Coupon, *errors.DomainError) {
	if len(Name) > 10 {
		return nil, errors.NewDomainError(errors.InvalidParameter, "クーポン名は10文字以内にしてください")
	}
	if ExpireAt.Before(time.Now()) {
		return nil, errors.NewDomainError(errors.InvalidParameter, "有効期限が現在より前にはできません")
	}
	return &Coupon{
		ID:                ID,
		Name:              Name,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		Notices:           Notices,
		User:              User,
		TargetStore:       TargetStore,
		CreateAt:          CreateAt,
		Status:            Status,
	}, nil
}

func RegenCoupon(
	ID uuid.UUID,
	Name string,
	CouponType CouponType,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	UsedAt *time.Time,
	User *User,
	TargetStore []*Store,
	CreateAt time.Time,
	Status CouponStatus,
) *Coupon {
	return &Coupon{
		ID:                ID,
		Name:              Name,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		Notices:           Notices,
		UsedAt:            UsedAt,
		User:              User,
		TargetStore:       TargetStore,
		CreateAt:          CreateAt,
		Status:            Status,
	}
}

func CreateStandardCoupon(
	User *User,
	TargetStore []*Store,
) (*Coupon, *errors.DomainError) {
	expireAtOneYear := time.Now().AddDate(1, 0, 0)

	return newCoupon(
		uuid.New(),
		"500円割引",
		Standard,
		500,
		expireAtOneYear,
		false,
		DefaultNotices,
		User,
		TargetStore,
		time.Now(),
		CouponIssued,
	)
}

func CreateCustomCoupon(
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	TargetStore []*Store,
) (*Coupon, *errors.DomainError) {
	if Name == "" {
		return nil, errors.NewDomainError(errors.InvalidParameter, "クーポン名が空です")
	}

	return newCoupon(
		uuid.New(),
		Name,
		Custom,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		nil,
		TargetStore,
		time.Now(),
		CouponDraft,
	)
}
func SaveCustomCoupon(
	DraftCoupon *Coupon,
) (*Coupon, *errors.DomainError) {

	return newCoupon(
		DraftCoupon.ID,
		DraftCoupon.Name,
		Custom,
		DraftCoupon.DiscountAmount,
		DraftCoupon.ExpireAt,
		DraftCoupon.IsCombinationable,
		DraftCoupon.Notices,
		nil,
		DraftCoupon.TargetStore,
		DraftCoupon.CreateAt,
		CouponSaved,
	)
}

func CreateBirthdayCoupon(
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	User *User,
	TargetStore []*Store,
) (*Coupon, *errors.DomainError) {
	return newCoupon(
		uuid.New(),
		"お誕生日",
		Birthday,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		DefaultNotices,
		User,
		TargetStore,
		time.Now(),
		CouponIssued,
	)
}

func UsedCoupon(coupon *Coupon) *Coupon {
	now := time.Now()

	return RegenCoupon(
		coupon.ID,
		coupon.Name,
		coupon.CouponType,
		coupon.DiscountAmount,
		coupon.ExpireAt,
		coupon.IsCombinationable,
		coupon.Notices,
		&now,
		coupon.User,
		coupon.TargetStore,
		coupon.CreateAt,
		CouponUsed,
	)
}
