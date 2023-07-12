package entity

import (
	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Name            string
	ZipCode         string
	Address         string
	Tel             string
	Parking         string
	AccessInfo      string
	IsActive        bool
	Stayable        bool //宿泊施設かどうか
	QRCode          uuid.UUID
	UnLimitedQRCode uuid.UUID
}

func CreateStore(
	Name string,
	Address string,
	Tel string,
	Parking string,
	AccessInfo string,
) *Store {
	return &Store{
		ID:              uuid.New(),
		Name:            Name,
		Address:         Address,
		Tel:             Tel,
		Parking:         Parking,
		AccessInfo:      AccessInfo,
		IsActive:        true,
		QRCode:          uuid.New(),
		UnLimitedQRCode: uuid.New(),
	}
}

func RegenStore(
	ID uuid.UUID,
	Name string,
	Address string,
	Tel string,
	Parking string,
	AccessInfo string,
	IsActive bool,
	QRCode uuid.UUID,
	UnLimitedQRCode uuid.UUID,
) *Store {
	return &Store{
		ID:              ID,
		Name:            Name,
		Address:         Address,
		Tel:             Tel,
		Parking:         Parking,
		AccessInfo:      AccessInfo,
		IsActive:        IsActive,
		QRCode:          QRCode,
		UnLimitedQRCode: UnLimitedQRCode,
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
