package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"

	"github.com/google/uuid"
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
	admin, err := models.FindAdmin(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	store := StoreModelToEntity(admin.R.BelongToStore)
	return AdminModelToEntity(admin, store, admin.R.Admin.Email), nil
}
func (pq *AdminQueryService) GetByMail(mail string) (*entity.Admin, error) {
	usermanager, err := models.UserManagers(models.UserManagerWhere.Email.EQ(mail)).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	admin := usermanager.R.AdminAdmin
	store := StoreModelToEntity(admin.R.BelongToStore)

	return AdminModelToEntity(admin, store, usermanager.Email), nil
}

func (pq *AdminQueryService) GetAll() ([]*entity.Admin, error) {
	admins, err := models.Admins(models.UserManagerWhere.IsAdmin.EQ(true)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	var result []*entity.Admin
	for _, admin := range admins {
		store := StoreModelToEntity(admin.R.BelongToStore)
		result = append(result, AdminModelToEntity(admin, store, admin.R.Admin.Email))
	}
	return result, nil
}

func AdminModelToEntity(model *models.Admin, store *entity.Store, email string) *entity.Admin {

	return entity.RegenAdmin(
		uuid.MustParse(model.AdminID),
		model.Name,
		email,
		model.IsActive,
		store,
	)
}
