package entity

import "github.com/google/uuid"

type UserOption struct {
	ID              uuid.UUID
	InnerNote       string
	IsBlackCustomer bool
}
type UserWithOption struct {
	UserOption
	User
}
