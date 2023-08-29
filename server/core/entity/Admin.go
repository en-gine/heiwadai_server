package entity

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID          uuid.UUID
	Name        string
	BelongStore Store
}

func CreateAdmin(
	Name string,
	BelongStore Store,
) *Admin {
	return &Admin{
		ID:          uuid.New(),
		Name:        Name,
		BelongStore: BelongStore,
	}
}
