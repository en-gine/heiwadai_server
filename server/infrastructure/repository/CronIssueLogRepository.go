package repository

import (
	"database/sql"

	"server/core/infra/repository"
	"server/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.ICronIssueLogRepository = &CronIssueLogRepository{}

type CronIssueLogRepository struct {
	db *sql.DB
}

func NewCronIssueLogRepository() *CronIssueLogRepository {
	db := InitDB()

	return &CronIssueLogRepository{
		db: db,
	}
}

func (pr *CronIssueLogRepository) Save(tx repository.ITransaction, cronName string, issueCount int, issueYear int, issueMonth int) error {
	cronLog := models.CronIssueLog{
		CronName:   cronName,
		IssueCount: issueCount,
		IssueYear:  issueYear,
		IssueMonth: issueMonth,
	}

	err := cronLog.Upsert(*tx.Context(), tx.Tran(), true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
