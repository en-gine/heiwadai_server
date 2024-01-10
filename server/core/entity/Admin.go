package entity

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID          uuid.UUID
	Name        string
	Mail        string
	IsActive    bool // 有効か(ログイン可能)どうか
	IsConfirmed bool // メール認証中かどうか
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
		IsConfirmed: false,
		BelongStore: BelongStore,
	}
}

func RegenAdmin(
	ID uuid.UUID,
	Name string,
	Mail string,
	IsActive bool,
	IsConfirmed bool,
	BelongStore *Store,
) *Admin {
	return &Admin{
		ID:          ID,
		Name:        Name,
		Mail:        Mail,
		IsActive:    IsActive,
		IsConfirmed: IsConfirmed,
		BelongStore: BelongStore,
	}
}
