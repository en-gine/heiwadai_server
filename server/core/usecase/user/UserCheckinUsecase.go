package user

import (
	"context"
	"time"

	"server/core/entity"
	"server/core/errors"
	queryService "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserCheckinUsecase struct {
	userQuery            queryService.IUserQueryService
	storeRepository      repository.IStoreRepository
	checkInRepository    repository.ICheckinRepository
	couponRepository     repository.ICouponRepository
	usercouponRepository repository.IUserCouponRepository
	usercouponQuery      queryService.IUserCouponQueryService
	storeQuery           queryService.IStoreQueryService
	checkinQuery         queryService.ICheckinQueryService
	couponQuery          queryService.ICouponQueryService
	transaction          repository.ITransaction
}

func NewUserCheckinUsecase(
	userQuery queryService.IUserQueryService,
	storeRepository repository.IStoreRepository,
	checkInRepository repository.ICheckinRepository,
	couponRepository repository.ICouponRepository,
	usercouponRepository repository.IUserCouponRepository,
	usercouponQuery queryService.IUserCouponQueryService,
	storeQuery queryService.IStoreQueryService,
	checkinQuery queryService.ICheckinQueryService,
	couponQuery queryService.ICouponQueryService,
	transaction repository.ITransaction,
) *UserCheckinUsecase {
	return &UserCheckinUsecase{
		userQuery:            userQuery,
		storeRepository:      storeRepository,
		checkInRepository:    checkInRepository,
		couponRepository:     couponRepository,
		usercouponRepository: usercouponRepository,
		usercouponQuery:      usercouponQuery,
		storeQuery:           storeQuery,
		checkinQuery:         checkinQuery,
		couponQuery:          couponQuery,
		transaction:          transaction,
	}
}

func (u *UserCheckinUsecase) GetStampCard(authID uuid.UUID) (*entity.StampCard, *errors.DomainError) {
	userCheckins, err := u.checkinQuery.GetActiveCheckin(authID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return entity.NewStampCard(userCheckins)
}

func (u *UserCheckinUsecase) Checkin(authID uuid.UUID, QrHash uuid.UUID) (*entity.UserAttachedCoupon, *errors.DomainError) {
	// チェックインによってクーポンが付与された場合クーポンを返す
	AuthUser, err := u.userQuery.GetByID(authID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if AuthUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}

	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	var checkInStore *entity.Store
	var isUnlimitQr bool
	lastCheckin, err := u.checkinQuery.GetLastStoreCheckin(authID, checkInStore.ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	for _, store := range allStores {
		// 通常のQRコードでは24時間以内にチェックインした店舗はチェックインできない
		if store.QRCode == QrHash {
			checkInStore = store
			isUnlimitQr = false
		}
		// 無制限のQRコード
		if store.UnLimitedQRCode == QrHash {
			checkInStore = store
			isUnlimitQr = true
		}
	}

	if checkInStore == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のQRコードの店舗が見つかりません。")
	}

	isSameStore := lastCheckin != nil && lastCheckin.Store.ID == checkInStore.ID

	if !isUnlimitQr && isSameStore && lastCheckin.CheckInAt.Add(24*time.Hour).After(time.Now()) {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "24時間以内にチェックインした店舗はチェックインできません。")
	}
	ctx := context.Background()
	err = u.transaction.Begin(ctx)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	newCheckin := entity.CreateCheckin(*checkInStore, *AuthUser)
	myCheckins, err := u.checkinQuery.GetActiveCheckin(authID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	var userAttachedCoupon *entity.UserAttachedCoupon
	if len(myCheckins) >= 5 {
		standardCoupon, domainErr := u.couponQuery.GetCouponByType(entity.CouponStandard)
		if domainErr != nil {
			u.transaction.Rollback()
			return nil, errors.NewDomainError(errors.QueryError, err.Error())
		}
		userAttachedCoupon := entity.CreateUserAttachedCoupon(authID, standardCoupon)

		err = u.usercouponRepository.Save(ctx, userAttachedCoupon)
		if err != nil {
			u.transaction.Rollback()
			return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
		}
	}

	err = u.checkInRepository.Save(ctx, newCheckin)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	err = u.transaction.Commit()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return userAttachedCoupon, nil
}
