package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserQueryService interface {
	GetByID(id uuid.UUID) (*entity.User, error)
	GetOptionByID(id uuid.UUID) (*entity.UserOption, error)
	GetByMail(mail entity.Mail) (*entity.User, error)
	GetMailOKUser(filterPrefectures *[]entity.Prefecture) ([]*entity.User, error)
	GetMailOKUserCount(filterPrefectures *[]entity.Prefecture) (*int, error)
	GetList(query *types.UserQuery, pager *types.PageQuery) ([]*entity.UserWichLastCheckin, *types.PageResponse, error)
	IsUnderRegister(mail entity.Mail) (bool, error)
}
