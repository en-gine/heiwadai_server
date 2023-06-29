package user

import (
	"errors"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserCouponUsecase struct {
	couponRepository repository.ICouponRepository
	couponQuery      queryservice.ICouponQueryService
}

func NewUserCouponUsecase(couponRepository repository.ICouponRepository, couponQuery queryservice.ICouponQueryService) *UserCouponUsecase {
	return &UserCouponUsecase{
		couponRepository: couponRepository,
		couponQuery:      couponQuery,
	}
}

func (u *UserCouponUsecase) GetMyList(AuthUser *entity.User) ([]*entity.Coupon, error) {

	coupons, err := u.couponQuery.GetActiveAll(AuthUser)

	return coupons, err
}

func (u *UserCouponUsecase) UseMyCoupon(AuthUser *entity.User, couponId uuid.UUID) error {
	coupon, err := u.couponQuery.GetById(couponId)

	if err != nil {
		return err
	}
	if coupon == nil {
		return errors.New("該当のクーポンIDが見つかりません。")
	}
	if AuthUser.ID != coupon.User.ID {
		return errors.New("ユーザー自身のクーポンではありません。")
	}
	if coupon.UsedAt != nil {
		return errors.New("既に使用済みのクーポンです。")
	}

	usedCoupon := entity.UsedCoupon(coupon)
	err = u.couponRepository.Save(usedCoupon)
	if err != nil {
		return err
	}

	return nil
}
