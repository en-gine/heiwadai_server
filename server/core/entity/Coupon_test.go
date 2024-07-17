package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCoupon(t *testing.T) {
	id := uuid.New()
	name := "Test"
	couponType := CouponStandard
	discountAmount := uint(500)
	expireAt := time.Now().Add(24 * time.Hour)
	isCombinationable := false
	notices := []string{"Test notice"}
	targetStore := []*Store{{ID: uuid.New(), Name: "Test Store"}}
	createAt := time.Now()
	status := CouponCreated

	coupon, err := newCoupon(id, name, couponType, discountAmount, expireAt, isCombinationable, notices, targetStore, createAt, status)
	fmt.Println(err)
	assert.Nil(t, err)
	assert.NotNil(t, coupon)
	assert.Equal(t, id, coupon.ID)
	assert.Equal(t, name, coupon.Name)
	assert.Equal(t, couponType, coupon.CouponType)
	assert.Equal(t, discountAmount, coupon.DiscountAmount)
	assert.Equal(t, expireAt, coupon.ExpireAt)
	assert.Equal(t, isCombinationable, coupon.IsCombinationable)
	assert.Equal(t, notices, coupon.Notices)
	assert.Equal(t, targetStore, coupon.TargetStore)
	assert.Equal(t, createAt, coupon.CreateAt)
	assert.Equal(t, status, coupon.Status)

	// Test invalid name
	_, err = newCoupon(id, "This is a very long coupon name", couponType, discountAmount, expireAt, isCombinationable, notices, targetStore, createAt, status)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "クーポン名は10文字以内にしてください")

	// Test invalid expireAt
	_, err = newCoupon(id, name, couponType, discountAmount, time.Now().Add(-24*time.Hour), isCombinationable, notices, targetStore, createAt, status)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "有効期限が現在より前にはできません")
}

func TestDefaultEmptyCustomCoupon(t *testing.T) {
	allStores := []*Store{{ID: uuid.New(), Name: "Test Store"}}
	coupon := DefaultEmptyCustomCoupon(allStores)

	assert.NotNil(t, coupon)
	assert.Equal(t, uuid.Nil, coupon.ID)
	assert.Empty(t, coupon.Name)
	assert.Equal(t, CouponCustom, coupon.CouponType)
	assert.Equal(t, uint(0), coupon.DiscountAmount)
	assert.WithinDuration(t, time.Now().AddDate(0, 1, 0), coupon.ExpireAt, time.Second)
	assert.False(t, coupon.IsCombinationable)
	assert.Equal(t, DefaultNotices, coupon.Notices)
	assert.Equal(t, allStores, coupon.TargetStore)
	assert.WithinDuration(t, time.Now(), coupon.CreateAt, time.Second)
	assert.Equal(t, CouponDraft, coupon.Status)
}

func TestCreateStandardCoupon(t *testing.T) {
	targetStore := []*Store{{ID: uuid.New(), Name: "Test Store"}}
	coupon, err := CreateStandardCoupon(targetStore)

	assert.Nil(t, err)
	assert.NotNil(t, coupon)
	assert.NotEqual(t, uuid.Nil, coupon.ID)
	assert.Equal(t, "500円", coupon.Name)
	assert.Equal(t, CouponStandard, coupon.CouponType)
	assert.Equal(t, uint(500), coupon.DiscountAmount)
	assert.WithinDuration(t, time.Now().AddDate(1, 0, 0), coupon.ExpireAt, time.Second)
	assert.False(t, coupon.IsCombinationable)
	assert.Equal(t, DefaultNotices, coupon.Notices)
	assert.Equal(t, targetStore, coupon.TargetStore)
	assert.WithinDuration(t, time.Now(), coupon.CreateAt, time.Second)
	assert.Equal(t, CouponIssued, coupon.Status)
}

func TestCreateBirthdayCoupon(t *testing.T) {
	targetStore := []*Store{{ID: uuid.New(), Name: "Test Store"}}
	coupon, err := CreateBirthdayCoupon(targetStore)

	assert.Nil(t, err)
	assert.NotNil(t, coupon)
	assert.NotEqual(t, uuid.Nil, coupon.ID)
	assert.Equal(t, "お誕生日", coupon.Name)
	assert.Equal(t, CouponBirthday, coupon.CouponType)
	assert.Equal(t, uint(500), coupon.DiscountAmount)
	assert.False(t, coupon.IsCombinationable)
	assert.Equal(t, DefaultNotices, coupon.Notices)
	assert.Equal(t, targetStore, coupon.TargetStore)
	assert.WithinDuration(t, time.Now(), coupon.CreateAt, time.Second)
	assert.Equal(t, CouponIssued, coupon.Status)

	// Check if ExpireAt is set to the last day of the current month
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	expectedExpireAt := time.Date(currentYear, currentMonth+1, 0, 23, 59, 59, 999999999, now.Location())
	assert.Equal(t, expectedExpireAt.Format("2006-01-02"), coupon.ExpireAt.Format("2006-01-02"))
}

func TestCreateCustomCoupon(t *testing.T) {
	name := "Custom"
	discountAmount := uint(1000)
	expireAt := time.Now().Add(30 * 24 * time.Hour)
	isCombinationable := true
	notices := []string{"Custom notice"}
	targetStore := []*Store{{ID: uuid.New(), Name: "Test Store"}}

	coupon, err := CreateCustomCoupon(name, discountAmount, expireAt, isCombinationable, notices, targetStore)

	assert.Nil(t, err)
	assert.NotNil(t, coupon)
	assert.NotEqual(t, uuid.Nil, coupon.ID)
	assert.Equal(t, name, coupon.Name)
	assert.Equal(t, CouponCustom, coupon.CouponType)
	assert.Equal(t, discountAmount, coupon.DiscountAmount)
	assert.Equal(t, expireAt, coupon.ExpireAt)
	assert.Equal(t, isCombinationable, coupon.IsCombinationable)
	assert.Equal(t, notices, coupon.Notices)
	assert.Equal(t, targetStore, coupon.TargetStore)
	assert.WithinDuration(t, time.Now(), coupon.CreateAt, time.Second)
	assert.Equal(t, CouponCreated, coupon.Status)

	// Test with empty name
	_, err = CreateCustomCoupon("", discountAmount, expireAt, isCombinationable, notices, targetStore)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "クーポン名が空です")
}

func TestCreateIssuedCoupon(t *testing.T) {
	originalCoupon, _ := CreateStandardCoupon([]*Store{{ID: uuid.New(), Name: "Test Store"}})
	count := 100

	issuedCoupon := CreateIssuedCoupon(originalCoupon, &count)

	assert.NotNil(t, issuedCoupon)
	assert.Equal(t, originalCoupon.ID, issuedCoupon.ID)
	assert.Equal(t, originalCoupon.Name, issuedCoupon.Name)
	assert.Equal(t, originalCoupon.CouponType, issuedCoupon.CouponType)
	assert.Equal(t, originalCoupon.DiscountAmount, issuedCoupon.DiscountAmount)
	assert.Equal(t, originalCoupon.ExpireAt, issuedCoupon.ExpireAt)
	assert.Equal(t, originalCoupon.IsCombinationable, issuedCoupon.IsCombinationable)
	assert.Equal(t, originalCoupon.Notices, issuedCoupon.Notices)
	assert.Equal(t, originalCoupon.TargetStore, issuedCoupon.TargetStore)
	assert.Equal(t, originalCoupon.CreateAt, issuedCoupon.CreateAt)
	assert.Equal(t, CouponIssued, issuedCoupon.Status)
	assert.Equal(t, &count, issuedCoupon.IssueCount)
	assert.NotNil(t, issuedCoupon.IssueAt)
	assert.WithinDuration(t, time.Now(), *issuedCoupon.IssueAt, time.Second)
}
