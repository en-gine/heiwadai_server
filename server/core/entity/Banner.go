package entity

import (
	"github.com/google/uuid"
)

type Banner struct {
	ID       uuid.UUID
	ImageURL string
	Url      string
}
