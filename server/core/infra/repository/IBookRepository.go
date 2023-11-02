package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IBookRepository interface {
	Save(book *entity.Booking) error
	Delete(bookID uuid.UUID) error
}
