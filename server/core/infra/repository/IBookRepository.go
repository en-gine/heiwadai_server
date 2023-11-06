package repository

import (
	"server/core/entity"

	"github.com/google/uuid"
)

type IBookRepository interface {
	Save(book *entity.Booking) error
	Delete(id uuid.UUID) error
}
type IBookAPIRepository interface {
	Reserve(bookData *entity.Booking) (*string, error)
	Cancel(bookData *entity.Booking) error
}
