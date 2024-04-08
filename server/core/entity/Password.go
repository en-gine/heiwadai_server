package entity

import (
	"server/core/errors"
	validator "server/validator"
)

type Password string

func NewPassword(password string) (*Password, *errors.DomainError) {
	v := validator.PasswordValidator{}

	_, err := v.IsValid(password)

	if err != nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}

	var pass = Password(password)
	return &pass, nil
}

func (p *Password) String() string {
	return string(*p)
}
