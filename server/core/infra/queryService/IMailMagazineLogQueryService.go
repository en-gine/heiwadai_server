package queryservice

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMailMagazineLogQueryService interface {
	GetUnsentTargetAllCount(mailMagazineID uuid.UUID) (int, error)
	GetUnsentTargetMails(mailMagazineID uuid.UUID, pager types.PageQuery) (*[]string, error)
	GetUserLogList(userID uuid.UUID, pager types.PageQuery) ([]*entity.MailMagazineLogWithTitle, *types.PageResponse, error)
}
