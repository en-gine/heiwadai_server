package entity

import (
	"server/core/errors"

	"github.com/google/uuid"
)

type Store struct {
	ID                uuid.UUID
	Name              string
	BranchName        *string
	ZipCode           string
	Address           string
	Tel               string
	SiteURL           string
	StampImageURL     string
	Stayable          bool
	IsActive          bool
	QRCode            uuid.UUID
	UnLimitedQRCode   uuid.UUID
	StayableStoreInfo *StayableStoreInfo
}

type StayableStoreInfo struct {
	Parking         string
	Latitude        float64
	Longitude       float64
	AccessInfo      string
	RestAPIURL      string
	BookingSystemID string
}

type StayableStore struct {
	*Store
	*StayableStoreInfo
}

func CreateStore(
	NewID uuid.UUID,
	Name string,
	BranchName *string,
	ZipCode string,
	Address string,
	Tel string,
	SiteURL string,
	StampImageURL string,
	Stayable bool,
	StayableStoreInfo *StayableStoreInfo,
) (*Store, *errors.DomainError) {
	if Stayable && StayableStoreInfo == nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, "宿泊可能な施設は宿泊施設情報を入力してください。")
	}

	return &Store{
		ID:                NewID,
		Name:              Name,
		BranchName:        BranchName,
		ZipCode:           ZipCode,
		Address:           Address,
		Tel:               Tel,
		SiteURL:           SiteURL,
		StampImageURL:     StampImageURL,
		Stayable:          Stayable,
		IsActive:          true,
		QRCode:            uuid.New(),
		UnLimitedQRCode:   uuid.New(),
		StayableStoreInfo: StayableStoreInfo,
	}, nil
}

func CreateStayableStoreInfo(
	Parking string,
	Latitude float64,
	Longitude float64,
	AccessInfo string,
	RestAPIURL string,
	BookingSystemID string,
) *StayableStoreInfo {
	return &StayableStoreInfo{
		Parking:         Parking,
		Longitude:       Longitude,
		Latitude:        Latitude,
		AccessInfo:      AccessInfo,
		RestAPIURL:      RestAPIURL,
		BookingSystemID: BookingSystemID,
	}
}

func CreateStayableStore(
	Store *Store,
	StayableInfo *StayableStoreInfo,
) (*StayableStore, *errors.DomainError) {
	if Store == nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, "Storeがnilです。")
	}
	if StayableInfo == nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, "StayableInfoがnilです。")
	}
	if !Store.Stayable {
		return nil, errors.NewDomainError(errors.InvalidParameter, "Stayableがtrueである必要があります。")
	}
	return &StayableStore{
		Store,
		StayableInfo,
	}, nil
}

func RegenStore(
	ID uuid.UUID,
	Name string,
	BranchName *string,
	ZipCode string,
	Address string,
	Tel string,
	SiteURL string,
	StampImageURL string,
	Stayable bool,
	IsActive bool,
	QRCode uuid.UUID,
	UnLimitedQRCode uuid.UUID,
	StayableStoreInfo *StayableStoreInfo,
) *Store {
	return &Store{
		ID:                ID,
		Name:              Name,
		BranchName:        BranchName,
		ZipCode:           ZipCode,
		Address:           Address,
		Tel:               Tel,
		SiteURL:           SiteURL,
		StampImageURL:     StampImageURL,
		Stayable:          Stayable,
		IsActive:          IsActive,
		QRCode:            QRCode,
		UnLimitedQRCode:   UnLimitedQRCode,
		StayableStoreInfo: StayableStoreInfo,
	}
}

func RegenStayableStoreInfo(
	Parking string,
	Latitude float64,
	Longitude float64,
	AccessInfo string,
	RestAPIURL string,
	BookingSystemID string,
) *StayableStoreInfo {
	return &StayableStoreInfo{
		Parking:         Parking,
		Latitude:        Latitude,
		Longitude:       Longitude,
		AccessInfo:      AccessInfo,
		RestAPIURL:      RestAPIURL,
		BookingSystemID: BookingSystemID,
	}
}

func RegenStayableStore(
	Store *Store,
	StayableInfo *StayableStoreInfo,
) *StayableStore {
	return &StayableStore{
		Store,
		StayableInfo,
	}
}
