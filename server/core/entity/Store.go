package entity

import (
	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Name            string
	Address         string
	IsActive        bool
	StayAble        bool //宿泊施設かどうか
	QrCode          uuid.UUID
	UnLimitedQrCode uuid.UUID
}

func CreateStore(
	Name string,
	Address string,
) *Store {
	return &Store{
		ID:              uuid.New(),
		Name:            Name,
		Address:         Address,
		IsActive:        true,
		QrCode:          uuid.New(),
		UnLimitedQrCode: uuid.New(),
	}
}

func RegenStore(
	ID uuid.UUID,
	Name string,
	Address string,
	IsActive bool,
	QrCode uuid.UUID,
	UnLimitedQrCode uuid.UUID,
) *Store {
	return &Store{
		ID:              ID,
		Name:            Name,
		Address:         Address,
		IsActive:        IsActive,
		QrCode:          QrCode,
		UnLimitedQrCode: UnLimitedQrCode,
	}
}
