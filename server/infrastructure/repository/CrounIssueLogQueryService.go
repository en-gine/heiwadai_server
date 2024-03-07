package repository

import (
	"context"
	"database/sql"

	queryservice "server/core/infra/queryService"
	"server/db/models"
	"server/infrastructure/logger"
)

var _ queryservice.ICronIssueLogQueryService = &CronIssueLogQueryService{}

type CronIssueLogQueryService struct {
	db *sql.DB
}

func NewCronIssueLogQueryService() *CronIssueLogQueryService {
	db := InitDB()

	return &CronIssueLogQueryService{
		db: db,
	}
}

func (pq *CronIssueLogQueryService) HasYearMonthLog(year int, month int) (bool, error) {
	log, err := models.CronIssueLogs(models.CronIssueLogWhere.IssueYear.EQ(year), models.CronIssueLogWhere.IssueMonth.EQ(month)).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logger.Error(err.Error())
		return false, err
	}
	if log == nil {
		return false, nil
	}
	return true, nil
}
