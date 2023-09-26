package repository

import (
	"github.com/google/uuid"
)

type IPlanRepository interface {
	Save(userID uuid.UUID, bookID string, StoreBookingSystemID string) error
	Delete(userID uuid.UUID, bookID string, StoreBookingSystemID string) error
}
