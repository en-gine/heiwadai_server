package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCheckin(t *testing.T) {
	store := Store{
		ID:   uuid.New(),
		Name: "Test Store",
	}
	user := User{
		ID:        uuid.New(),
		FirstName: "Test",
		LastName:  "User",
	}

	checkin := CreateCheckin(store, user)

	assert.NotNil(t, checkin)
	assert.NotEqual(t, uuid.Nil, checkin.ID)
	assert.Equal(t, &store, checkin.Store)
	assert.Equal(t, &user, checkin.User)
	assert.WithinDuration(t, time.Now(), checkin.CheckInAt, time.Second)
	assert.False(t, checkin.Archive)
}

func TestRegenCheckin(t *testing.T) {
	id := uuid.New()
	store := &Store{
		ID:   uuid.New(),
		Name: "Regen Store",
	}
	user := &User{
		ID:        uuid.New(),
		FirstName: "Regen",
		LastName:  "User",
	}
	checkInAt := time.Now().Add(-1 * time.Hour)
	archive := true

	checkin := RegenCheckin(id, store, user, checkInAt, archive)

	assert.NotNil(t, checkin)
	assert.Equal(t, id, checkin.ID)
	assert.Equal(t, store, checkin.Store)
	assert.Equal(t, user, checkin.User)
	assert.Equal(t, checkInAt, checkin.CheckInAt)
	assert.True(t, checkin.Archive)
}
