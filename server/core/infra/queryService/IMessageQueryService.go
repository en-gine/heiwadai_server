package queryservice

import (
	"time"

	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMessageQueryService interface {
	GetByID(id uuid.UUID) (*entity.Message, error)
	GetMessagesAfter(lastCreateAt *time.Time) ([]*entity.Message, error)
	GetAll(pager *types.PageQuery) ([]*entity.Message, error)
}
