package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserAttachedCouponUsecase struct {
	usercouponRepository repository.IUserCouponRepository
	usercouponQuery      queryservice.IUserCouponQueryService
}

func NewUserAttachedCouponUsecase(usercouponRepository repository.IUserCouponRepository, usercouponQuery queryservice.IUserCouponQueryService) *UserAttachedCouponUsecase {
	return &UserAttachedCouponUsecase{
		usercouponRepository: usercouponRepository,
		usercouponQuery:      usercouponQuery,
	}
}

func (u *UserAttachedCouponUsecase) GetMyList(AuthUser *entity.User) ([]*entity.UserAttachedCoupon, *errors.DomainError) {

	usercoupons, err := u.usercouponQuery.GetActiveAll(AuthUser)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return usercoupons, nil
}

func (u *UserAttachedCouponUsecase) UseMyCoupon(AuthUser *entity.User, couponID uuid.UUID) error {
	coupon, err := u.usercouponQuery.GetByID(AuthUser, couponID)

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

	usedCoupon := entity.UseUserAttachedCoupon(
		AuthUser,
		coupon.Coupon,
	)
	err = u.usercouponRepository.Save(usedCoupon)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}
