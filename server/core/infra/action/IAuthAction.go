package action

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/types"

	"github.com/google/uuid"
)

type UserType string

const (
	UserTypeAdmin UserType = "admin"
	UserTypeUser  UserType = "user"
)

type UserAuth struct {
	UserID   uuid.UUID
	UserType UserType
	Token    *types.Token
}

type UserInfo struct {
	UserID   uuid.UUID
	UserType UserType
	Mail     string
}

func (ut UserType) String() string {
	return string(ut)
}

type IAuthAction interface {
	SignUp(email entity.Mail, password entity.Password) (*uuid.UUID, error)
	SignIn(email entity.Mail, password string) (*types.Token, *errors.DomainError, error)
	SignOut(token string) error
	Refresh(token string, refreshToken string) (*UserAuth, error)
	GetUserInfo(token string) (*UserInfo, error)
	ResetPasswordMail(email entity.Mail) (*errors.DomainError, error)
	UpdatePassword(password entity.Password, token string) (*errors.DomainError, error)
	InviteUserByEmail(email entity.Mail) (*uuid.UUID, *errors.DomainError, error)
	ReInviteUserByEmail(email entity.Mail) (*errors.DomainError, error)
	UpdateEmail(email entity.Mail, token string) error
}
