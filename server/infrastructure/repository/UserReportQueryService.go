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
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

var _ queryservice.IUserReportQueryService = &UserReportQueryService{}

type UserReportQueryService struct {
	db *sql.DB
}

func NewUserReportQueryService() *UserReportQueryService {
	db := InitDB()

	return &UserReportQueryService{
		db: db,
	}
}

func (pq *UserReportQueryService) GetByID(id uuid.UUID) (*entity.UserReport, error) {
	rpt, err := models.FindUserReport(context.Background(), pq.db, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if rpt == nil {
		return nil, nil
	}
	return UserReportModelToEntity(rpt), nil
}

func (pq *UserReportQueryService) GetAll(pager *types.PageQuery) ([]*entity.UserReport, error) {
	msgs, err := models.UserReports(qm.Limit(pager.Limit()), qm.Offset(pager.Offset()), qm.OrderBy(models.UserReportColumns.CreateAt+" DESC")).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	var result []*entity.UserReport
	for _, rpt := range msgs {
		result = append(result, UserReportModelToEntity(rpt))
	}
	return result, nil
}

func UserReportModelToEntity(rpt *models.UserReport) *entity.UserReport {
	return &entity.UserReport{
		ID:       uuid.MustParse(rpt.ID),
		Title:    rpt.Title,
		Content:  rpt.Content,
		UserID:   uuid.MustParse(rpt.UserID),
		UserName: rpt.UserName,
	}
}
