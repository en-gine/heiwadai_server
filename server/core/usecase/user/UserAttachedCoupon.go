package user

import (
	"context"

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

func (u *UserAttachedCouponUsecase) GetByID(AuthUserID uuid.UUID, couponID uuid.UUID) (*entity.UserAttachedCoupon, *errors.DomainError) {
	usercoupon, err := u.usercouponQuery.GetByID(AuthUserID, couponID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return usercoupon, nil
}

func (u *UserAttachedCouponUsecase) GetMyList(AuthUserID uuid.UUID) ([]*entity.UserAttachedCoupon, *errors.DomainError) {
	usercoupons, err := u.usercouponQuery.GetActiveAll(AuthUserID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return usercoupons, nil
}

func (u *UserAttachedCouponUsecase) UseMyCoupon(AuthUserID uuid.UUID, couponID uuid.UUID) *errors.DomainError {
	coupon, err := u.usercouponQuery.GetByID(AuthUserID, couponID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if AuthUserID != coupon.UserID {
		return errors.NewDomainError(errors.InvalidParameter, "ユーザー自身のクーポンではありません。")
	}
	if coupon.Status != entity.CouponIssued {
		return errors.NewDomainError(errors.UnPemitedOperation, "発行済ステータスのクーポンではありません。")
	}

	if coupon.UsedAt != nil {
		return errors.NewDomainError(errors.UnPemitedOperation, "クーポンはすでに使用済みです。")
	}

	usedCoupon := entity.UseUserAttachedCoupon(
		AuthUserID,
		coupon.Coupon,
	)

	err = u.usercouponRepository.Save(context.Background(), usedCoupon)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}
