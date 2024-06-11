package user

import (
	"context"
	"reflect"
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
	userCheckins, err := u.checkinQuery.GetMyActiveCheckin(authID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return entity.NewStampCard(userCheckins)
}

func (u *UserCheckinUsecase) Checkin(authID uuid.UUID, QrHash uuid.UUID) (*entity.StampCard, *entity.UserAttachedCoupon, *errors.DomainError) {
	// チェックインによってクーポンが付与された場合クーポンを返す
	AuthUser, err := u.userQuery.GetByID(authID)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if AuthUser == nil {
		return nil, nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のユーザーが見つかりません。")
	}

	qrStore, err := u.storeQuery.GetStoreByQrCode(QrHash)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	unlimitQrStore, err := u.storeQuery.GetStoreByUnlimitQrCode(QrHash)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if qrStore == nil && unlimitQrStore == nil {
		return nil, nil, errors.NewDomainError(errors.QueryDataNotFoundError, "該当のQRコードの店舗が見つかりません。")
	}

	var checkInStore *entity.Store
	var isUnlimitQr bool
	if qrStore != nil {
		checkInStore = qrStore
		isUnlimitQr = false
	}
	if unlimitQrStore != nil {
		checkInStore = unlimitQrStore
		isUnlimitQr = true
	}

	lastCheckin, err := u.checkinQuery.GetMyLastStoreCheckin(authID, checkInStore.ID)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	isNil := lastCheckin == nil || reflect.ValueOf(lastCheckin).IsNil()

	var isSameStore bool = false
	if !isNil {
		isSameStore = lastCheckin.Store.ID == checkInStore.ID
	}

	var isAfter24Hours bool = false
	if isSameStore {
		isAfter24Hours = lastCheckin.CheckInAt.Add(24 * time.Hour).After(time.Now())
	}

	if !isUnlimitQr && isSameStore && isAfter24Hours {
		return nil, nil, errors.NewDomainError(errors.UnPemitedOperation, "24時間以内にチェックインした店舗はチェックインできません。")
	}
	ctx := context.Background()
	err = u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	newCheckin := entity.CreateCheckin(*checkInStore, *AuthUser)
	myCheckins, err := u.checkinQuery.GetMyActiveCheckin(authID)
	if err != nil {
		u.transaction.Rollback()
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	err = u.checkInRepository.Save(u.transaction, newCheckin)
	if err != nil {
		u.transaction.Rollback()
		return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	var userAttachedCoupon *entity.UserAttachedCoupon
	if len(myCheckins)+1 >= 5 {
		allStores, err := u.storeQuery.GetActiveAll()
		if err != nil {
			u.transaction.Rollback()
			return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
		}
		standardCoupon, domainErr := entity.CreateStandardCoupon(allStores)
		if domainErr != nil {
			u.transaction.Rollback()
			return nil, nil, domainErr
		}

		err = u.couponRepository.Save(u.transaction, standardCoupon)
		if err != nil {
			u.transaction.Rollback()
			return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
		}

		userAttachedCoupon = entity.CreateUserAttachedCoupon(authID, standardCoupon)

		err = u.usercouponRepository.Save(u.transaction, userAttachedCoupon)
		if err != nil {
			u.transaction.Rollback()
			return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
		}
		var count int = 1
		issuedCoupon := entity.CreateIssuedCoupon(standardCoupon, &count)
		err = u.couponRepository.Save(u.transaction, issuedCoupon)
		if err != nil {
			u.transaction.Rollback()
			return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
		}

		err = u.checkInRepository.BulkArchive(u.transaction, authID)
		if err != nil {
			u.transaction.Rollback()
			return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
		}
	}

	err = u.transaction.Commit()
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	NewStampCard, domainErr := u.GetStampCard(authID)
	if domainErr != nil {
		return nil, nil, domainErr
	}
	return NewStampCard, userAttachedCoupon, nil
}
