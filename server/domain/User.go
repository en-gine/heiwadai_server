package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID          uuid.UUID
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   *string
	BirthDate     time.Time
	ZipCode       *string
	Prefecture    string
	City          *string
	Address       *string
	Tel           *string
	Mail          string
	AcceptMail    bool
}

func (u *User) New(
	UUID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture string,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool,
) *User {
	return &User{
		UUID:          UUID,
		FirstName:     FirstName,
		LastName:      LastName,
		FirstNameKana: FirstNameKana,
		LastNameKana:  LastNameKana,
		CompanyName:   CompanyName,
		BirthDate:     BirthDate,
		ZipCode:       ZipCode,
		Prefecture:    Prefecture,
		City:          City,
		Address:       Address,
		Tel:           Tel,
		Mail:          Mail,
		AcceptMail:    AcceptMail,
	}
}
