package entity

import (
	"server/core/errors"
	validator "server/validator"
)

type Mail string

func NewMail(password string) (*Mail, *errors.DomainError) {
	v := validator.EmailValidator{}

	_, err := v.IsValid(password)

	if err != nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}

	var pass = Mail(password)
	return &pass, nil
}

func (p *Mail) String() string {
	return string(*p)
}
