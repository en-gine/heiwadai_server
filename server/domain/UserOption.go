package domain

import (
	"math/big"
)

type UserOption struct {
	userId          big.Int
	innerNote       string
	isBlackCustomer bool
}

func (u *UserOption) NewUserOption(
	userId big.Int,
	innerNote string,
	isBlackCustomer bool,
) *UserOption {
	return &UserOption{
		userId:          userId,
		innerNote:       innerNote,
		isBlackCustomer: isBlackCustomer,
	}
}
