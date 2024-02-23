package admin

import (
	"context"
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type AdminCouponUsecase struct {
	couponRepository     repository.ICouponRepository
	couponQuery          queryservice.ICouponQueryService
	userCouponQuery      queryservice.IUserCouponQueryService
	usercouponRepository repository.IUserCouponRepository
	storeQuery           queryservice.IStoreQueryService
	transaction          repository.ITransaction
}

func NewAdminCouponUsecase(couponRepository repository.ICouponRepository, couponQuery queryservice.ICouponQueryService,
	userCouponQuery queryservice.IUserCouponQueryService, usercouponRepository repository.IUserCouponRepository, storeQuery queryservice.IStoreQueryService,
	transaction repository.ITransaction,
) *AdminCouponUsecase {
	return &AdminCouponUsecase{
		couponRepository:     couponRepository,
		couponQuery:          couponQuery,
		userCouponQuery:      userCouponQuery,
		usercouponRepository: usercouponRepository,
		storeQuery:           storeQuery,
		transaction:          transaction,
	}
}

func (u *AdminCouponUsecase) CreateStandardCoupon() *errors.DomainError {
	// Seederで叩く想定。デフォルトのクーポンをDB生成＆保持
	store, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	standard, domainErr := entity.CreateStandardCoupon(store)
	if err != nil {
		return domainErr
	}
	ctx := context.Background()
	tx := u.transaction
	err = tx.Begin(ctx)

	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	err = u.couponRepository.Save(tx, standard)
	if err != nil {
		tx.Rollback()
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AdminCouponUsecase) GetUsersCouponList(UserID uuid.UUID, pager *types.PageQuery) ([]*entity.UserAttachedCoupon, *types.PageResponse, *errors.DomainError) {
	coupons, pageRes, err := u.userCouponQuery.GetAll(UserID, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return coupons, pageRes, nil
}

func (u *AdminCouponUsecase) DefaultEmptyCustomCoupon() (*entity.Coupon, *errors.DomainError) {
	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	stores := make([]*entity.Store, 0)
	stores = append(stores, allStores...)
	entity := entity.DefaultEmptyCustomCoupon(stores)
	return entity, nil
}

func (u *AdminCouponUsecase) CreateCustomCoupon(
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	TargetStore []*entity.Store,
) (*entity.Coupon, *errors.DomainError) {
	customCoupon, domainErr := entity.CreateCustomCoupon(
		Name,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		TargetStore,
	)
	if domainErr != nil {
		return nil, domainErr
	}
	ctx := context.Background()
	err := u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.couponRepository.Save(u.transaction, customCoupon)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.transaction.Commit()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return customCoupon, nil
}

func (u *AdminCouponUsecase) SaveCustomCoupon(
	couponID uuid.UUID,
	Name string,
	DiscountAmount uint,
	ExpireAt time.Time,
	IsCombinationable bool,
	Notices []string,
	TargetStore []*entity.Store,
) *errors.DomainError {
	coupon, err := u.couponQuery.GetByID(couponID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponCreated {
		return errors.NewDomainError(errors.UnPemitedOperation, "データ保存済の状態のクーポンではありません。")
	}

	coupon, domainErr := entity.SaveCustomCoupon(
		coupon.ID,
		Name,
		DiscountAmount,
		ExpireAt,
		IsCombinationable,
		Notices,
		TargetStore,
		coupon.CreateAt,
	)

	if domainErr != nil {
		return domainErr
	}
	ctx := context.Background()
	err = u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	err = u.couponRepository.Save(u.transaction, coupon)
	if err != nil {
		u.transaction.Rollback()
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.transaction.Commit()
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}

func (u *AdminCouponUsecase) GetCustomCouponByID(couponID uuid.UUID) (*entity.Coupon, *errors.DomainError) {
	coupon, err := u.couponQuery.GetByID(couponID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.CouponType != entity.CouponCustom {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	return coupon, nil
}

func (u *AdminCouponUsecase) GetCustomCouponList(pager *types.PageQuery) ([]*entity.Coupon, *types.PageResponse, *errors.DomainError) {
	coupons, pageRes, err := u.couponQuery.GetCouponListByType(entity.CouponCustom, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return coupons, pageRes, nil
}

func (u *AdminCouponUsecase) AttachCustomCouponToAllUser(couponID uuid.UUID) (*int, *errors.DomainError) {
	coupon, err := u.couponQuery.GetByID(couponID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if coupon == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のクーポンIDが見つかりません。")
	}
	if coupon.Status != entity.CouponCreated {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "保存済ステータスのクーポンではありません。")
	}

	ctx := context.Background()
	err = u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	count, err := u.usercouponRepository.IssueAll(u.transaction, coupon, nil)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}
	if count == 0 {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.ActionError, "クーポンの発行に失敗しました。")
	}
	issuedCoupon := entity.CreateIssuedCoupon(coupon, &count)
	err = u.couponRepository.Save(u.transaction, issuedCoupon)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.transaction.Commit()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return &count, nil
}

func (u *AdminCouponUsecase) BulkAttachBirthdayCoupon(birthMonth int) (*int, *errors.DomainError) {
	ctx := context.Background()
	err := u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	birthdayCoupon, domainErr := entity.CreateBirthdayCoupon(allStores)
	if err != nil {
		u.transaction.Rollback()
		return nil, domainErr
	}

	err = u.couponRepository.Save(u.transaction, birthdayCoupon)
	if domainErr != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	count, err := u.usercouponRepository.IssueAll(u.transaction, birthdayCoupon, &birthMonth)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}

	issuedCoupon := entity.CreateIssuedCoupon(birthdayCoupon, &count)
	err = u.couponRepository.Save(u.transaction, issuedCoupon)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.transaction.Commit()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return &count, nil
}
