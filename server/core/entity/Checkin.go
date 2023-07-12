package entity

import (
	"time"

	"github.com/google/uuid"
)

type Checkin struct {
	ID uuid.UUID
	*Store
	*User
	CheckInAt time.Time
	Archive   bool
}

func CreateCheckin(
	store Store,
	user User,
) *Checkin {
	return &Checkin{
		ID:        uuid.New(),
		Store:     &store,
		User:      &user,
		CheckInAt: time.Now(),
		Archive:   false,
	}
}

func RegenCheckin(
	id uuid.UUID,
	store *Store,
	user *User,
	at time.Time,
	archive bool,
) *Checkin {
	return &Checkin{
		ID:        id,
		Store:     store,
		User:      user,
		CheckInAt: at,
		Archive:   archive,
	}
}
