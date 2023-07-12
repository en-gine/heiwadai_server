package user

import (
	"server/core/entity"
	"server/core/errors"
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

func (u *UserCouponUsecase) GetMyList(AuthUser *entity.User) ([]*entity.Coupon, *errors.DomainError) {

	coupons, err := u.couponQuery.GetActiveAll(AuthUser)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return coupons, nil
}

func (u *UserCouponUsecase) UseMyCoupon(AuthUser *entity.User, couponId uuid.UUID) *errors.DomainError {
	coupon, err := u.couponQuery.GetById(couponId)

	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if AuthUser.ID != coupon.User.ID {
		return errors.NewDomainError(errors.InvalidParameter, "ユーザー自身のクーポンではありません。")
	}
	if coupon.Status != entity.CouponIssued {
		return errors.NewDomainError(errors.UnPemitedOperation, "発行済ステータスのクーポンではありません。")
	}

	if coupon.UsedAt != nil {
		return errors.NewDomainError(errors.UnPemitedOperation, "クーポンはすでに使用済みです。")
	}

	usedCoupon := entity.UsedCoupon(coupon)
	err = u.couponRepository.Save(usedCoupon)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}
