package queryservice

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IBookQueryService interface {
	GetMyBooking(userID uuid.UUID) ([]*entity.Booking, error)
	GetBookRequestNumber() (*string, error)
}
