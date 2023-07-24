package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
)

type AdminCouponUsecase struct {
	couponRepository repository.ICouponRepository
	couponQuery      queryservice.ICouponQueryService
	userCouponQuery  queryservice.IUserCouponQueryService
	couponAction     action.IAttachCouponAction
	storeQuery       queryservice.IStoreQueryService
}

func NewAdminCouponUsecase(couponRepository repository.ICouponRepository, couponQuery queryservice.ICouponQueryService,
	userCouponQuery queryservice.IUserCouponQueryService, couponAction action.IAttachCouponAction, storeQuery queryservice.IStoreQueryService) *AdminCouponUsecase {
	return &AdminCouponUsecase{
		couponRepository: couponRepository,
		couponQuery:      couponQuery,
		couponAction:     couponAction,
		storeQuery:       storeQuery,
	}
}

// Seederで叩く想定。デフォルトのクーポンをDB生成＆保持
func (u *AdminCouponUsecase) CreateDefaultCoupon() *errors.DomainError {
	store, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	standard, domainErr := entity.CreateStandardCoupon(store)
	if err != nil {
		return domainErr
	}
	err = u.couponRepository.Save(standard)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AdminCouponUsecase) GetUsersCouponList(User *entity.User) ([]*entity.UserAttachedCoupon, *errors.DomainError) {

	coupons, err := u.userCouponQuery.GetActiveAll(User)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return coupons, nil
}

func (u *AdminCouponUsecase) CreateCustomCoupon(
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
) (*entity.Coupon, *errors.DomainError) {

	stores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	customCoupon, domainErr := entity.CreateCustomCoupon(
		Name,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		stores,
	)
	if domainErr != nil {
		return nil, domainErr
	}
	err = u.couponRepository.Save(customCoupon)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return customCoupon, nil
}
func (u *AdminCouponUsecase) SaveCustomCoupon(couponId uuid.UUID) error {
	coupon, err := u.couponQuery.GetByID(couponId)

	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponDraft {
		return errors.NewDomainError(errors.UnPemitedOperation, "下書き状態のクーポンではありません。")
	}

	saveCoupon, domainErr := entity.SaveCustomCoupon(coupon)
	if domainErr != nil {
		return domainErr
	}

	err = u.couponRepository.Save(saveCoupon)

	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}
func (u *AdminCouponUsecase) AttachCustomCouponToAllUser(couponId uuid.UUID) (*int, *errors.DomainError) {
	coupon, err := u.couponQuery.GetByID(couponId)

	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponSaved {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "保存済ステータスのクーポンではありません。")
	}

	count, err := u.couponAction.Issue(coupon)
	if err != nil {
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}
	if count == 0 {
		return nil, errors.NewDomainError(errors.ActionError, "クーポンの発行に失敗しました。")
	}
	return &count, nil
}
