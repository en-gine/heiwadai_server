package queryservice

import (
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMailMagazineLogQueryService interface {
	GetUnsentTargetAllCount(mailMagazineID uuid.UUID) (int, error)
	GetUnsentTargetMails(mailMagazineID uuid.UUID, pager types.PageQuery) (*[]string, error)
}
