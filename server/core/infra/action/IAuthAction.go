package action

import (
	"server/core/entity"
	"server/core/infra/types"
)

type IAuthAction interface {
	SignUp(email string, password string, UserData *entity.User) (*entity.User, error)
	SignIn(email string, password string) (*types.Token, error)
	GetUser(token *types.Token) (*entity.User, error)
	Refresh(token *types.Token) (*types.Token, error)
}
