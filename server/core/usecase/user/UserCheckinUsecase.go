package user

import (
	"errors"
	"server/core/entity"
	queryService "server/core/infra/queryService"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
)

type UserCheckinUsecase struct {
	storeRepository   repository.IStoreRepository
	checkInRepository repository.ICheckinRepository
	couponRepository  repository.ICouponRepository
	storeQuery        queryService.IStoreQueryService
	checkinQuery      queryService.ICheckinQueryService
	couponQuery       queryService.ICouponQueryService
	transaction       repository.ITransaction
}

func NewUserCheckinUsecase(
	storeRepository repository.IStoreRepository,
	checkInRepository repository.ICheckinRepository,
	couponRepository repository.ICouponRepository,
	storeQuery queryService.IStoreQueryService,
	checkinQuery queryService.ICheckinQueryService,
	couponQuery queryService.ICouponQueryService,
	transaction repository.ITransaction,

) *UserCheckinUsecase {
	return &UserCheckinUsecase{
		storeRepository:   storeRepository,
		checkInRepository: checkInRepository,
		couponRepository:  couponRepository,
		storeQuery:        storeQuery,
		checkinQuery:      checkinQuery,
		couponQuery:       couponQuery,
		transaction:       transaction,
	}
}

func (u *UserCheckinUsecase) GetStampCard(user *entity.User) (*entity.StampCard, error) {
	userCheckins, err := u.checkinQuery.GetActiveCheckin(user)
	if err != nil {
		return nil, err
	}
	return entity.NewStampCard(userCheckins)
}

// チェックインによってクーポンが付与された場合クーポンを返す
func (u *UserCheckinUsecase) Checkin(AuthUser *entity.User, QrHash uuid.UUID) (*entity.Coupon, error) {

	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, err
	}
	var checkInStore *entity.Store
	var isUnlimitQr bool
	lastCheckin, err := u.checkinQuery.GetLastStoreCheckin(AuthUser, *checkInStore)

	if err != nil {
		return nil, err
	}

	for _, store := range allStores {
		//通常のQRコードでは24時間以内にチェックインした店舗はチェックインできない
		if store.QrCode == QrHash {
			checkInStore = store
			isUnlimitQr = false
		}
		//無制限のQRコード
		if store.UnLimitedQrCode == QrHash {
			checkInStore = store
			isUnlimitQr = true
		}
	}

	if checkInStore == nil {
		return nil, errors.New("該当のQRコードの店舗が見つかりません。")
	}

	isSameStore := lastCheckin != nil && lastCheckin.Store.ID == checkInStore.ID

	if !isUnlimitQr && isSameStore && lastCheckin.CheckInAt.Add(24*time.Hour).After(time.Now()) {
		return nil, errors.New("24時間以内にチェックインした店舗はチェックインできません。")
	}

	u.transaction.Begin()
	newCheckin := entity.CreateCheckin(*checkInStore, *AuthUser)
	myCheckins, err := u.checkinQuery.GetActiveCheckin(AuthUser)
	if err != nil {
		u.transaction.Rollback()
		return nil, err
	}
	var newCoupon *entity.Coupon
	if len(myCheckins) >= 5 {
		newCoupon, err = entity.CreateStandardCoupon(AuthUser, allStores)
		if err != nil {
			u.transaction.Rollback()
			return nil, err
		}
		err = u.couponRepository.Save(newCoupon)
	}

	err = u.checkInRepository.Save(newCheckin)
	if err != nil {
		u.transaction.Rollback()
		return nil, err
	}
	u.transaction.Commit()

	return newCoupon, nil
}
