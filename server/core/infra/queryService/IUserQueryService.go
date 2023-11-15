package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserQueryService interface {
	GetByID(id uuid.UUID) (*entity.User, error)
	GetByMail(mail string) (*entity.User, error)
	GetMailOKUser(filterPrefectures *[]entity.Prefecture) ([]*entity.User, error)
	GetMailOKUserCount(filterPrefectures *[]entity.Prefecture) (*int, error)
	GetAll(query *types.UserQuery, pager *types.PageQuery) ([]*entity.User, error)
}
