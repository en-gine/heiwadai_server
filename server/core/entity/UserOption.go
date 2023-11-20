package entity

import "github.com/google/uuid"

type UserOption struct {
	ID              uuid.UUID
	InnerNote       string
	IsBlackCustomer bool
}
type UserWithOption struct {
	User       *User
	UserOption *UserOption
}

func CreateUserOption(
	ID uuid.UUID,
	InnerNote string,
	IsBlackCustomer bool,
) *UserOption {
	return &UserOption{
		ID:              ID,
		InnerNote:       InnerNote,
		IsBlackCustomer: IsBlackCustomer,
	}
}

func RegenUserOption(
	ID uuid.UUID,
	InnerNote string,
	IsBlackCustomer bool,
) *UserOption {
	return &UserOption{
		ID:              ID,
		InnerNote:       InnerNote,
		IsBlackCustomer: IsBlackCustomer,
	}
}
