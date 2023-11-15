package repository

import (
	"server/core/entity"
	"server/core/infra/queryService/types"

	"github.com/google/uuid"
)

type IMailMagazineLogRepository interface {
	BulkCopyToLogAsUnsent(magazineID uuid.UUID, filterPref *[]entity.Prefecture) error
	BulkMarkAsSent(magazineID uuid.UUID, pager types.PageQuery) error
}
