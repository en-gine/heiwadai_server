package entity

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID       uuid.UUID
	Name     string
	BelongTo Store
}

func CreateAdmin(
	Name string,
	BelongTo Store,
) *Admin {
	return &Admin{
		ID:       uuid.New(),
		Name:     Name,
		BelongTo: BelongTo,
	}
}
