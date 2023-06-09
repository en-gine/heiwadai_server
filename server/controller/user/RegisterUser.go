package user

import (
	"server/service/user"

	"github.com/Songmu/go-httpdate"
)

type RegisterUserRequest struct {
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	FirstNameKana string  `json:"firstNameKana"`
	LastNameKana  string  `json:"lastNameKana"`
	CompanyName   *string `json:"companyName"`
	BirthDate     string  `json:"birthDate"`
	ZipCode       *string `json:"zipCode"`
	Prefecture    string  `json:"prefecture"`
	City          *string `json:"city"`
	Address       *string `json:"address"`
	Tel           *string `json:"tel"`
	Mail          string  `json:"mail"`
	AcceptMail    bool    `json:"acceptMail"`
}

type RegisterUserResponse struct {
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	FirstNameKana string  `json:"firstNameKana"`
	LastNameKana  string  `json:"lastNameKana"`
	CompanyName   *string `json:"companyName"`
	BirthDate     string  `json:"birthDate"`
	ZipCode       *string `json:"zipCode"`
	Prefecture    string  `json:"prefecture"`
	City          *string `json:"city"`
	Address       *string `json:"address"`
	Tel           *string `json:"tel"`
	Mail          string  `json:"mail"`
	AcceptMail    bool    `json:"acceptMail"`
}

func Controller(request RegisterUserRequest) (*RegisterUserResponse, error) {
	birth, err := httpdate.Str2Time(request.BirthDate, nil)
	if err != nil {
		return nil, err
	}

	service := user.RegisterUserService{}
	user, err := service.Exec(
		request.FirstName,
		request.LastName,
		request.FirstNameKana,
		request.LastNameKana,
		request.CompanyName,
		birth,
		request.ZipCode,
		request.Prefecture,
		request.City,
		request.Address,
		request.Tel,
		request.Mail,
		request.AcceptMail,
	)
	if err != nil {
		return nil, err
	}

	return &RegisterUserResponse{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FirstNameKana: user.FirstNameKana,
		LastNameKana:  user.LastNameKana,
		CompanyName:   user.CompanyName,
		BirthDate:     user.BirthDate.String(),
		ZipCode:       user.ZipCode,
		Prefecture:    user.Prefecture,
		City:          user.City,
		Address:       user.Address,
		Tel:           user.Tel,
		Mail:          user.Mail,
		AcceptMail:    user.AcceptMail,
	}, err
}
