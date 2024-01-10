package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IAdminQueryService = &AdminQueryService{}

type AdminQueryService struct {
	db *sql.DB
}

func NewAdminQueryService() *AdminQueryService {
	db := InitDB()

	return &AdminQueryService{
		db: db,
	}
}

func (pq *AdminQueryService) GetByID(id uuid.UUID) (*entity.Admin, error) {
	admin, err := models.Admins(qm.Load(models.AdminRels.BelongToStore), qm.Load(models.AdminRels.Admin), models.AdminWhere.AdminID.EQ(id.String())).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if admin == nil {
		return nil, nil
	}
	store := StoreModelToEntity(admin.R.BelongToStore, nil)
	return AdminModelToEntity(admin, store, admin.R.Admin.Email), nil
}

func (pq *AdminQueryService) GetByMail(mail string) (*entity.Admin, error) {
	usermanager, err := models.UserManagers(
		models.UserManagerWhere.Email.EQ(mail),
		qm.Load(models.UserManagerRels.AdminAdmin),
		models.UserManagerWhere.IsAdmin.EQ(true)).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	if usermanager == nil {
		return nil, nil
	}

	admin := usermanager.R.AdminAdmin
	if admin == nil {
		return nil, nil
	}
	modelStore, err := models.Stores(models.StoreWhere.ID.EQ(admin.BelongTo)).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	if modelStore == nil {
		return nil, nil
	}
	store := StoreModelToEntity(modelStore, nil)

	return AdminModelToEntity(admin, store, usermanager.Email), nil
}

func (pq *AdminQueryService) GetAll() ([]*entity.Admin, error) {
	admins, err := models.Admins(qm.Load(models.AdminRels.Admin), qm.Load(models.AdminRels.BelongToStore)).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if admins == nil {
		return nil, nil
	}
	var result []*entity.Admin
	for _, admin := range admins {
		store := StoreModelToEntity(admin.R.BelongToStore, nil)
		result = append(result, AdminModelToEntity(admin, store, admin.R.Admin.Email))
	}
	return result, nil
}

func GetAdminIsConfirmed(adminID uuid.UUID) (bool, error) {
	db := InitDB()
	var confirmedAt *string
	err := db.QueryRow("SELECT email_confirmed_at FROM auth.users WHERE id = $1", adminID.String()).Scan(&confirmedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logger.Error(err.Error())
		return false, err
	}
	if confirmedAt == nil {
		return false, nil
	}
	return true, nil
}

func AdminModelToEntity(model *models.Admin, store *entity.Store, email string) *entity.Admin {
	isConfirmed, err := GetAdminIsConfirmed(uuid.MustParse(model.AdminID))
	if err != nil {
		return nil
	}
	return entity.RegenAdmin(
		uuid.MustParse(model.AdminID),
		model.Name,
		email,
		model.IsActive,
		isConfirmed,
		store,
	)
}
