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
	userLoginLogs := models.UserLoginLogs(models.UserLoginLogWhere.UserID.EQ(userID.String()), qm.OrderBy("login_at DESC"), qm.Limit(pager.Limit()), qm.Offset(pager.Offset()))

	userLoginLogModels, err := userLoginLogs.All(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}
	count, err := models.MailMagazines().Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}

	var result []*entity.UserLoginLog
	for _, userLoginLog := range userLoginLogModels {
		result = append(result, UserLoginLogModelToEntity(userLoginLog))
	}

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
