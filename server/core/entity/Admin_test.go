package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdmin(t *testing.T) {
	store := &Store{} // テスト用のStoreを作成
	name := "Test Admin"
	mail := "admin@example.com"

	admin := CreateAdmin(name, mail, store)

	assert.NotNil(t, admin)
	assert.NotEqual(t, uuid.Nil, admin.ID)
	assert.Equal(t, name, admin.Name)
	assert.Equal(t, mail, admin.Mail)
	assert.True(t, admin.IsActive)
	assert.False(t, admin.IsConfirmed)
	assert.Equal(t, store, admin.BelongStore)
}

func TestRegenAdmin(t *testing.T) {
	id := uuid.New()
	name := "Regen Admin"
	mail := "regen@example.com"
	isActive := false
	isConfirmed := true
	store := &Store{} // テスト用のStoreを作成

	admin := RegenAdmin(id, name, mail, isActive, isConfirmed, store)

	assert.NotNil(t, admin)
	assert.Equal(t, id, admin.ID)
	assert.Equal(t, name, admin.Name)
	assert.Equal(t, mail, admin.Mail)
	assert.Equal(t, isActive, admin.IsActive)
	assert.Equal(t, isConfirmed, admin.IsConfirmed)
	assert.Equal(t, store, admin.BelongStore)
}
