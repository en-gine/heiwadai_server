package domain

import (
	"github.com/google/uuid"
)

type UserOption struct {
	UserUUId        uuid.UUID
	InnerNote       string
	IsBlackCustomer bool
}

func (u *UserOption) New(
	UserUUId uuid.UUID,
	innerNote string,
	isBlackCustomer bool,
) *UserOption {
	return &UserOption{
		UserUUId:        UserUUId,
		InnerNote:       innerNote,
		IsBlackCustomer: isBlackCustomer,
	}
}
