package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IMailMagazineLogQueryService = &MailMagazineLogQueryService{}

type MailMagazineLogQueryService struct {
	db *sql.DB
}

func NewMailMagazineLogQueryService() *MailMagazineLogQueryService {
	db := InitDB()

	return &MailMagazineLogQueryService{
		db: db,
	}
}

func (pq *MailMagazineLogQueryService) GetUnsentTargetAllCount(mailMagazineID uuid.UUID) (int, error) {
	count, err := models.MailMagazineLogs(models.MailMagazineLogWhere.MailMagazineID.EQ(mailMagazineID.String())).Count(context.Background(), pq.db)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (pq *MailMagazineLogQueryService) GetUserLogList(userID uuid.UUID, pager types.PageQuery) ([]*entity.MailMagazineLogWithTitle, *types.PageResponse, error) {
	logs, err := models.MailMagazineLogs(models.MailMagazineLogWhere.UserID.EQ(userID.String()),
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset()),
		qm.Load(models.MailMagazineLogRels.MailMagazine),
	).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}
	var result []*entity.MailMagazineLogWithTitle
	for _, log := range logs {
		result = append(result, MailMagazineLogModelToEntity(log))
	}
	count, err := models.MailMagazineLogs(
		models.MailMagazineLogWhere.UserID.EQ(userID.String()),
	).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	pageResponse := types.NewPageResponse(&pager, int(count))
	return result, pageResponse, err
}

func (pq *MailMagazineLogQueryService) GetUnsentTargetMails(mailMagazineID uuid.UUID, pager types.PageQuery) (*[]string, error) {
	mails, err := models.MailMagazineLogs(qm.Select(models.MailMagazineLogColumns.Email), models.MailMagazineLogWhere.MailMagazineID.EQ(mailMagazineID.String()), qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	mailStrs := make([]string, len(mails))

	for i, mail := range mails {
		mailStrs[i] = mail.Email
	}
	return &mailStrs, nil
}

func MailMagazineLogModelToEntity(model *models.MailMagazineLog) *entity.MailMagazineLogWithTitle {
	return &entity.MailMagazineLogWithTitle{
		Log: &entity.MailMagazineLog{
			MailMagazineID: uuid.MustParse(model.MailMagazineID),
			UserID:         uuid.MustParse(model.UserID),
			Email:          model.Email,
			SentAt:         &model.SentAt.Time,
		},
		Title: model.R.MailMagazine.Title,
	}
}
