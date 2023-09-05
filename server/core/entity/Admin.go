package entity

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID          uuid.UUID
	Name        string
	Mail        string
	IsActive    bool
	BelongStore *Store
}

func CreateAdmin(
	Name string,
	Mail string,
	BelongStore *Store,
) *Admin {
	return &Admin{
		ID:          uuid.New(),
		Name:        Name,
		Mail:        Mail,
		IsActive:    true,
		BelongStore: BelongStore,
	}
}

func RegenAdmin(
	ID uuid.UUID,
	Name string,
	Mail string,
	IsActive bool,
	BelongStore *Store,
) *Admin {
	return &Admin{
		ID:          ID,
		Name:        Name,
		Mail:        Mail,
		IsActive:    IsActive,
		BelongStore: BelongStore,
	}
}
