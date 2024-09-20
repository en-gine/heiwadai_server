package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	storeRepository repository.IStoreRepository
	storeQuery      queryservice.IStoreQueryService
	fileUploader    action.IFileAction
}

func NewStoreUsecase(storeRepository repository.IStoreRepository, storeQuery queryservice.IStoreQueryService, fileUploader action.IFileAction) *StoreUsecase {
	return &StoreUsecase{
		storeRepository: storeRepository,
		storeQuery:      storeQuery,
		fileUploader:    fileUploader,
	}
}

func (u *StoreUsecase) GetStoreByID(storeID uuid.UUID) (*entity.Store, *errors.DomainError) {
	store, err := u.storeQuery.GetByID(storeID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return store, nil
}

func (u *StoreUsecase) GetActiveList() ([]*entity.Store, *errors.DomainError) {
	stores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return stores, nil
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
	StampImageData string,
	Stayable bool,
	Parking *string,
	Latitude *float64,
	Longitude *float64,
	AccessInfo *string,
	RestAPIURL *string,
	BookingSystemID *string,
	BookingSystemLoginID *string,
	BookingSystemLoginPassword *string,
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
		if BookingSystemLoginID == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemLoginIDは必須です。")
		}
		if BookingSystemLoginPassword == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemLoginPasswordは必須です。")
		}
		stayableInfo = entity.CreateStayableStoreInfo(
			*Parking,
			*Latitude,
			*Longitude,
			*AccessInfo,
			*RestAPIURL,
			*BookingSystemID,
			*BookingSystemLoginID,
			*BookingSystemLoginPassword,
		)
	}

	newID := uuid.New()
	StampImageURL, err := u.fileUploader.Upload(&StampImageData, newID.String())
	if err != nil {
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}
	store, domainErr := entity.CreateStore(
		newID,
		Name,
		BranchName,
		ZipCode,
		Address,
		Tel,
		SiteURL,
		*StampImageURL,
		Stayable,
		stayableInfo,
	)
	if domainErr != nil {
		return nil, domainErr
	}

	err = u.storeRepository.Save(store, stayableInfo)
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
	StampImageData string,
	Stayable bool,
	Parking *string,
	Latitude *float64,
	Longitude *float64,
	AccessInfo *string,
	RestAPIURL *string,
	BookingSystemID *string,
	BookingSystemLoginID *string,
	BookingSystemLoginPassword *string,
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
		if BookingSystemLoginID == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemLoginIDは必須です。")
		}
		if BookingSystemLoginPassword == nil {
			return nil, errors.NewDomainError(errors.InvalidParameter, "BookingSystemLoginPasswordは必須です。")
		}
		stayableInfo = entity.CreateStayableStoreInfo(
			*Parking,
			*Latitude,
			*Longitude,
			*AccessInfo,
			*RestAPIURL,
			*BookingSystemID,
			*BookingSystemLoginID,
			*BookingSystemLoginPassword,
		)
	}
	var StampImageURL *string
	if StampImageData != "" {
		StampImageURL, err = u.fileUploader.Upload(&StampImageData, storeID.String())
	} else {
		StampImageURL = &existStore.StampImageURL
	}

	if StampImageURL == nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, "スタンプ画像は必須です。")
	}

	if err != nil {
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}
	updateStore := entity.RegenStore(
		storeID,
		Name,
		BranchName,
		ZipCode,
		Address,
		Tel,
		SiteURL,
		*StampImageURL,
		Stayable,
		IsActive,
		QRCode,
		UnLimitedQRCode,
		stayableInfo,
	)

	err = u.storeRepository.Save(updateStore, stayableInfo)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateStore, nil
}

func (u *StoreUsecase) RegenQR(storeID uuid.UUID) (*uuid.UUID, *errors.DomainError) {
	store, err := u.storeQuery.GetByID(storeID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if store == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の店舗が見つかりません")
	}

	qr, err := u.storeRepository.RegenQR(store.ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return qr, nil
}

func (u *StoreUsecase) RegenUnlimitQR(storeID uuid.UUID) (*uuid.UUID, *errors.DomainError) {
	store, err := u.storeQuery.GetByID(storeID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if store == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の店舗が見つかりません")
	}

	qr, err := u.storeRepository.RegenUnlimitQR(store.ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return qr, nil
}
