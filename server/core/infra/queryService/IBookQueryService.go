package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IBookQueryService interface {
	GetByID(bookID uuid.UUID) (*entity.Booking, error)
	GetMyBooking(userID uuid.UUID) ([]*entity.Booking, error)
	GetBookRequestDataID() (*string, error)
}
