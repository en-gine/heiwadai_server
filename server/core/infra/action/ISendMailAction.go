package action

import (
	"server/core/entity"
)

type IMailAction interface {
	SendAll(users []*entity.User) error
}
