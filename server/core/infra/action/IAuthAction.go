package action

import (
	"server/core/entity"
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
	SignUp(email string, password entity.Password) (*uuid.UUID, error)
	SignIn(email string, password entity.Password) (*types.Token, error)
	SignOut(token string) error
	Refresh(token string, refreshToken string) (*UserAuth, error)
	GetUserInfo(token string) (*UserInfo, error)
	ResetPasswordMail(email string) error
	UpdatePassword(password entity.Password, token string) error
	InviteUserByEmail(mail string) (*uuid.UUID, error)
	UpdateEmail(email string, token string) error
}
