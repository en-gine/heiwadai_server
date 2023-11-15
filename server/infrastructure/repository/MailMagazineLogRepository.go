package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/entity"
	"server/core/infra/queryService/types"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ repository.IMailMagazineLogRepository = &MailMagazineLogRepository{}

type MailMagazineLogRepository struct {
	db *sql.DB
}

func NewMailMagazineLogRepository() *MailMagazineLogRepository {
	db := InitDB()

	return &MailMagazineLogRepository{
		db: db,
	}
}

func (pq *MailMagazineLogRepository) BulkCopyToLogAsUnsent(magazineID uuid.UUID, filterPref *[]entity.Prefecture) error {
	mailUserWhere := GetMailUserWhereMods(filterPref)
	queryMods := []qm.QueryMod{
		qm.SQL("INSERT INTO ?", models.TableNames.MailMagazineLog),
		qm.SQL("(?, ?, ?)", models.MailMagazineLogColumns.MailMagazineID, models.MailMagazineLogColumns.UserID, models.UserManagerColumns.Email),
		qm.Select(magazineID.String(), models.UserDatumColumns.UserID),
		qm.From(models.TableNames.UserData),
		qm.InnerJoin(models.TableNames.UserManager+" on ? = ?", models.UserManagerColumns.ID, models.UserDatumColumns.UserID),
	}
	queryMods = append(queryMods, mailUserWhere...)

	_, err := models.NewQuery(
		queryMods...,
	).Exec(pq.db)

	return err
}

func (pq *MailMagazineLogRepository) BulkMarkAsSent(magazineID uuid.UUID, pager types.PageQuery) error {
	ctx := context.Background()
	logs, err := models.MailMagazineLogs(models.MailMagazineLogWhere.MailMagazineID.EQ(magazineID.String()), qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(ctx, pq.db)
	if err != nil {
		return err
	}
	_, err = logs.UpdateAll(ctx, pq.db, models.M{models.MailMagazineLogColumns.SentAt: time.Now()})
	return err
}
