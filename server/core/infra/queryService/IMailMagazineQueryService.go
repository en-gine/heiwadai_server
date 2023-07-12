package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMailMagazineQueryService interface {
	GetById(id uuid.UUID) (*entity.MailMagazine, error)
	GetAll(pager *types.PageQuery) ([]*entity.MailMagazine, error) //statusは不問
}
