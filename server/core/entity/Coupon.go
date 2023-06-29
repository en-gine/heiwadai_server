package entity

import (
	"errors"
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
	OverView          string     //概要
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

func newCoupon(
	ID uuid.UUID,
	Name string,
	CouponType CouponType,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	OverView string,
	User *User,
	TargetStore []*Store,
) (*Coupon, error) {
	if len(Name) > 10 {
		return nil, errors.New("クーポン名は10文字以内にしてください")
	}
	if ExpireAt.Before(time.Now()) {
		return nil, errors.New("有効期限が現在より前にはできません")
	}
	return &Coupon{
		ID:                ID,
		Name:              Name,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		OverView:          OverView,
		User:              User,
		TargetStore:       TargetStore,
	}, nil
}

func StoredCoupon(
	ID uuid.UUID,
	Name string,
	CouponType CouponType,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	OverView string,
	UsedAt *time.Time,
	User *User,
	TargetStore []*Store,
) *Coupon {
	return &Coupon{
		ID:                ID,
		Name:              Name,
		DiscountAmount:    DiscountAmount,
		ExpireAt:          ExpireAt,
		IsCombinationable: IsCombinationable,
		OverView:          OverView,
		UsedAt:            UsedAt,
		User:              User,
		TargetStore:       TargetStore,
	}
}

func CreateStandardCoupon(
	User *User,
	TargetStore []*Store,
) (*Coupon, error) {
	expireAtOneYear := time.Now().AddDate(1, 0, 0)

	return newCoupon(
		uuid.New(),
		"500円割引",
		Standard,
		500,
		expireAtOneYear,
		false,
		"",
		User,
		TargetStore,
	)
}

func CreateCustomCoupon(
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	OverView string,
	User *User,
	TargetStore []*Store,
) (*Coupon, error) {
	if Name == "" {
		return nil, errors.New("クーポン名が空です")
	}

	return newCoupon(
		uuid.New(),
		Name,
		Custom,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		OverView,
		User,
		TargetStore,
	)
}

func CreateBirthdayCoupon(
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	OverView string,
	User *User,
	TargetStore []*Store,
) (*Coupon, error) {
	return newCoupon(
		uuid.New(),
		"お誕生日",
		Birthday,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		OverView,
		User,
		TargetStore,
	)
}

func UsedCoupon(coupon *Coupon) *Coupon {
	now := time.Now()
	coupon.UsedAt = &now
	return StoredCoupon(
		coupon.ID,
		coupon.Name,
		coupon.CouponType,
		coupon.DiscountAmount,
		coupon.ExpireAt,
		coupon.IsCombinationable,
		coupon.OverView,
		coupon.UsedAt,
		coupon.User,
		coupon.TargetStore,
	)
}
