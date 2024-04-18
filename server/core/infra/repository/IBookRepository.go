package repository

import (
	"server/core/entity"
	"server/core/errors"

	"github.com/google/uuid"
)

type IBookRepository interface {
	Save(book *entity.Booking) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
}
type IBookAPIRepository interface {
	Reserve(bookData *entity.Booking) (*string, *errors.DomainError, error)
	Cancel(bookData *entity.Booking, newDataID string) (*errors.DomainError, error)
}
