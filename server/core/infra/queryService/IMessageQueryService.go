package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMessageQueryService interface {
	GetByID(id uuid.UUID) (*entity.Message, error)
	GetMessagesAfter(id *uuid.UUID) ([]*entity.Message, error)
	GetAll(pager *types.PageQuery) ([]*entity.Message, error)
}
