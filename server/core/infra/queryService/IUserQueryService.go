package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IUserQueryService interface {
	GetByID(id uuid.UUID) (*entity.User, error)
	GetByMail(mail string) (*entity.User, error)
	GetUserByPrefecture([]*entity.Prefecture) ([]*entity.User, error)
	GetAll(pager *types.PageQuery) ([]*entity.User, error)
}
