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
	StayableStoreInfo *StayableStoreInfo
	IsActive          bool
	QRCode            uuid.UUID
	UnLimitedQRCode   uuid.UUID
}

type StayableStoreInfo struct {
	Parking         string
	Latitude        float64
	Longitude       float64
	AccessInfo      string
	RestAPIURL      string
	BookingSystemID string
}

func CreateStore(
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
		return nil, errors.NewDomainError(errors.InvalidParameter, "Stayableがtrueの場合、StayableStoreInfoは必須です。")
	}
	return &Store{
		ID:                uuid.New(),
		Name:              Name,
		BranchName:        BranchName,
		ZipCode:           ZipCode,
		Address:           Address,
		Tel:               Tel,
		SiteURL:           SiteURL,
		StampImageURL:     StampImageURL,
		Stayable:          Stayable,
		StayableStoreInfo: StayableStoreInfo,
		IsActive:          true,
		QRCode:            uuid.New(),
		UnLimitedQRCode:   uuid.New(),
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
	StayableStoreInfo *StayableStoreInfo,
	IsActive bool,
	QRCode uuid.UUID,
	UnLimitedQRCode uuid.UUID,
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
		StayableStoreInfo: StayableStoreInfo,
		IsActive:          IsActive,
		QRCode:            QRCode,
		UnLimitedQRCode:   UnLimitedQRCode,
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

func (s *Store) StayableStore(
	Stores []*Store,
) []*Store {
	var stayableStore []*Store
	for _, store := range Stores {
		if store.Stayable {
			stayableStore = append(stayableStore, store)
		}
	}
	return stayableStore
}
