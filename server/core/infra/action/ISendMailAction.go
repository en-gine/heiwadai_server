package action

import (
	"server/core/entity"
)

type IMailAction interface {
	SendAll(prefecture []*entity.Prefecture) error
}
