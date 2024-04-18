package queryservice

import (
	"server/core/entity"
	"time"

	"github.com/google/uuid"
)

type IBookQueryService interface {
	GetByID(bookID uuid.UUID) (*entity.Booking, error)
	GetMyBooking(userID uuid.UUID, stayFromAfterAt time.Time) ([]*entity.Booking, error)
	GenerateBookDataID() (*string, error)
}
