package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IUserLoginLogRepository = &UserLoginLogRepository{}

type UserLoginLogRepository struct {
	db *sql.DB
}

func NewUserLoginLogRepository() *UserLoginLogRepository {
	db := InitDB()

	return &UserLoginLogRepository{
		db: db,
	}
}

func (pr *UserLoginLogRepository) Save(loginLog *entity.UserLoginLog) error {
	loginlog := models.UserLoginLog{
		UserID:    loginLog.UserID.String(),
		LoginAt:   loginLog.LoginAt,
		RemoteIP:  loginLog.RemoteID,
		UserAgent: loginLog.UserAgent,
	}

	err := loginlog.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
