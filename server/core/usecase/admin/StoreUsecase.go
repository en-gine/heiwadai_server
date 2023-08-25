package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	storeRepository repository.IStoreRepository
	storeQuery      queryservice.IStoreQueryService
}

func NewStoreUsecase(storeRepository repository.IStoreRepository, storeQuery queryservice.IStoreQueryService) *StoreUsecase {
	return &StoreUsecase{
		storeRepository: storeRepository,
		storeQuery:      storeQuery,
	}
}

func (u *StoreUsecase) GetList() ([]*entity.Store, *errors.DomainError) {

	stores, err := u.storeQuery.GetAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return stores, nil
}

func (u *StoreUsecase) Create(
	Name string,
	BranchName *string,
	ZipCode string,
	Address string,
	Tel string,
	SiteURL string,
	StampImageURL string,
	Stayable bool,
	Parking *string,
	Latitude *float64,
	Longitude *float64,
	AccessInfo *string,
	RestAPIURL *string,
	BookingSystemID *string,
) (*entity.Store, *errors.DomainError) {

	var stayableInfo *entity.StayableStoreInfo
	if Stayable {
		if Parking == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Parkingは必須です。")
		}
		if Latitude == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Latitudeは必須です。")
		}
		if Longitude == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Longitudeは必須です。")
		}
		if AccessInfo == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "AccessInfoは必須です。")
		}
		if RestAPIURL == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "AccessInfoは必須です。")
		}
		if BookingSystemID == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemIDは必須です。")
		}
		stayableInfo = entity.CreateStayableStoreInfo(
			*Parking,
			*Latitude,
			*Longitude,
			*AccessInfo,
			*RestAPIURL,
			*BookingSystemID,
		)
	}

	store, domainErr := entity.CreateStore(
		Name,
		BranchName,
		ZipCode,
		Address,
		Tel,
		SiteURL,
		StampImageURL,
		Stayable,
		stayableInfo,
	)
	if domainErr != nil {
		return nil, domainErr
	}
	err := u.storeRepository.Save(store)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return store, nil
}

func (u *StoreUsecase) Update(
	storeID uuid.UUID,
	Name string,
	BranchName *string,
	ZipCode string,
	Address string,
	Tel string,
	SiteURL string,
	StampImageURL string,
	Stayable bool,
	Parking *string,
	Latitude *float64,
	Longitude *float64,
	AccessInfo *string,
	RestAPIURL *string,
	BookingSystemID *string,
	IsActive bool,
	QRCode uuid.UUID,
	UnLimitedQRCode uuid.UUID,
) (*entity.Store, *errors.DomainError) {

	existStore, err := u.storeQuery.GetByID(storeID)

	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, "店舗のデータ問合せに失敗しました。")
	}

	if existStore == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の店舗が見つかりません")
	}
	var stayableInfo *entity.StayableStoreInfo
	if Stayable {
		if Parking == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Parkingは必須です。")
		}
		if Latitude == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Latitudeは必須です。")
		}
		if Longitude == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "Longitudeは必須です。")
		}
		if AccessInfo == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "AccessInfoは必須です。")
		}
		if RestAPIURL == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "AccessInfoは必須です。")
		}
		if BookingSystemID == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemIDは必須です。")
		}
		stayableInfo = entity.CreateStayableStoreInfo(
			*Parking,
			*Latitude,
			*Longitude,
			*AccessInfo,
			*RestAPIURL,
			*BookingSystemID,
		)
	}
	updateStore := entity.RegenStore(
		storeID,
		Name,
		BranchName,
		ZipCode,
		Address,
		Tel,
		SiteURL,
		StampImageURL,
		Stayable,
		stayableInfo,
		IsActive,
		QRCode,
		UnLimitedQRCode,
	)

	err = u.storeRepository.Save(updateStore)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateStore, nil
}
