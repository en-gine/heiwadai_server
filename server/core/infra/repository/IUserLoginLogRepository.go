package repository

import (
	"server/core/entity"
)

type IUserLoginLogRepository interface {
	Save(loginLog *entity.UserLoginLog) error
}
