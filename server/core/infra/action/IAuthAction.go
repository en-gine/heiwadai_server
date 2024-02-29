package action

import (
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

func (ut UserType) String() string {
	return string(ut)
}

type IAuthAction interface {
	SignUp(email string, password string, userType UserType) (*uuid.UUID, error)
	SignIn(email string, password string) (*types.Token, error)
	SignOut(token string) error
	Refresh(token string, refreshToken string) (*UserAuth, error)
	// GetUserID(token string) (*uuid.UUID, *UserType, error)
	ResetPasswordMail(email string) error
	UpdatePassword(password string, token string) error
	InviteUserByEmail(mail string) (*uuid.UUID, error)
	UpdateEmail(email string, token string) error
}
