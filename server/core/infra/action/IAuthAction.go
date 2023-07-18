package action

import (
	"server/core/infra/types"

	"github.com/google/uuid"
)

type IAuthAction interface {
	SignUp(email string, password string) error
	SignIn(email string, password string) (*types.Token, error)
	ResetPasswordMail(email string) error
	UpdatePassword(password string, token string) error
	InviteUserByEmail(mail string) (uuid.UUID, error)
	UpdateEmail(email string, token string) error
}
