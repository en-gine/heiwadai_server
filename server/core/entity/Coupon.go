package entity

import (
	"time"
	"unicode/utf8"

	"server/core/errors"

	"github.com/google/uuid"
)

type Coupon struct {
	ID                uuid.UUID
	Name              string
	CouponType        CouponType
	DiscountAmount    uint      // 割引額
	ExpireAt          time.Time // 有効期限
	IsCombinationable bool      // 併用可能
	Notices           []string  // 注意事項
	TargetStore       []*Store  // 対象店舗
	CreateAt          time.Time
	Status            CouponStatus
}

var DefaultNotices = []string{"クーポンは併用できません", "ランチではお使いになれません"}

type CouponType int

const (
	CouponStandard CouponType = iota
	CouponCustom
	CouponBirthday
)

func (c CouponType) String() string {
	switch c {
	case CouponStandard:
		return "Standard"
	case CouponCustom:
		return "Custom"
	case CouponBirthday:
		return "Birthday"
	default:
		return "Unknown"
	}
}

func (c CouponType) ToInt() int {
	return int(c)
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
	TargetStore []*Store,
	CreateAt time.Time,
	Status CouponStatus,
) (*Coupon, *errors.DomainError) {
	if utf8.RuneCountInString(Name) > 10 {
		return nil, errors.NewDomainError(errors.InvalidParameter, "クーポン名は10文字以内にしてください")
	}
	if ExpireAt.Before(time.Now()) {
		return nil, errors.NewDomainError(errors.InvalidParameter, "有効期限が現在より前にはできません")
	}
	return &Coupon{
		ID:                ID,
		Name:              Name,
		CouponType:        CouponType,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		Notices:           Notices,
		TargetStore:       TargetStore,
		CreateAt:          CreateAt,
		Status:            Status,
	}, nil
}

func DefaultEmptyCustomCoupon(allStores []*Store) *Coupon {
	return &Coupon{
		ID:                uuid.Nil,
		Name:              "",
		CouponType:        CouponCustom,
		DiscountAmount:    0,
		ExpireAt:          time.Now().AddDate(0, 1, 0),
		IsCombinationable: true,
		Notices:           DefaultNotices,
		TargetStore:       allStores,
		CreateAt:          time.Now(),
		Status:            CouponDraft,
	}
}

func CreateStandardCoupon(
	TargetStore []*Store,
) (*Coupon, *errors.DomainError) {
	expireAtOneYear := time.Now().AddDate(1, 0, 0)

	return newCoupon(
		uuid.New(),
		"500円割引",
		CouponStandard,
		500,
		expireAtOneYear,
		false,
		DefaultNotices,
		TargetStore,
		time.Now(),
		CouponIssued,
	)
}

func CreateBirthdayCoupon(
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	TargetStore []*Store,
) (*Coupon, *errors.DomainError) {
	return newCoupon(
		uuid.New(),
		"お誕生日",
		CouponBirthday,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		DefaultNotices,
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
		CouponCustom,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		TargetStore,
		time.Now(),
		CouponDraft,
	)
}

func SaveCustomCoupon(
	ID uuid.UUID,
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	TargetStore []*Store,
	CreateAt time.Time,
) (*Coupon, *errors.DomainError) {
	return RegenCoupon(
		ID,
		Name,
		CouponCustom,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		TargetStore,
		CreateAt,
		CouponSaved,
	), nil
}

func RegenCoupon(
	ID uuid.UUID,
	Name string,
	CouponType CouponType,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	TargetStore []*Store,
	CreateAt time.Time,
	Status CouponStatus,
) *Coupon {
	return &Coupon{
		ID:                ID,
		Name:              Name,
		CouponType:        CouponType,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		Notices:           Notices,
		TargetStore:       TargetStore,
		CreateAt:          CreateAt,
		Status:            Status,
	}
}
