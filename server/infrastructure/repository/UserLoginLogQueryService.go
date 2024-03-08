package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IUserLoginLogQueryService = &UserLoginLogQueryService{}

type UserLoginLogQueryService struct {
	db *sql.DB
}

func NewUserLoginLogQueryService() *UserLoginLogQueryService {
	db := InitDB()

	return &UserLoginLogQueryService{
		db: db,
	}
}

func (pq *UserLoginLogQueryService) GetList(userID uuid.UUID, pager *types.PageQuery) ([]*entity.UserLoginLog, *types.PageResponse, error) {
	logs, err := models.UserLoginLogs(models.UserLoginLogWhere.UserID.EQ(userID.String()),
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	var result []*entity.UserLoginLog
	for _, log := range logs {
		result = append(result, UserLoginLogModelToEntity(log))
	}

	count, err := models.UserLoginLogs(
		models.UserLoginLogWhere.UserID.EQ(userID.String()),
	).Count(context.Background(), pq.db)

	if err != nil {
		return nil, nil, err
	}

	pageResponse := types.NewPageResponse(pager, int(count))
	return result, pageResponse, nil
}

func UserLoginLogModelToEntity(model *models.UserLoginLog) *entity.UserLoginLog {
	return entity.RegenUserLoginLog(
		uuid.MustParse(model.UserID),
		model.LoginAt,
		model.RemoteIP,
		model.UserAgent,
	)
}
