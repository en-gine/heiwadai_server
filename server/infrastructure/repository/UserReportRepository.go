package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	repository "server/core/infra/repository"
	"server/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IUserReportRepository = &UserReportRepository{}

type UserReportRepository struct {
	db *sql.DB
}

func NewUserReportRepository() *UserReportRepository {
	db := InitDB()

	return &UserReportRepository{
		db: db,
	}
}

func (pq *UserReportRepository) Save(report *entity.UserReport) error {
	msg := models.UserReport{
		ID:       report.ID.String(),
		Title:    report.Title,
		Content:  report.Content,
		UserID:   report.UserID.String(),
		UserName: report.UserName,
	}

	err := msg.Upsert(context.Background(), pq.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
