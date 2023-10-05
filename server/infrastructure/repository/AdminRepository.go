package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IAdminRepository = &AdminRepository{}

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository() *AdminRepository {
	db := InitDB()

	return &AdminRepository{
		db: db,
	}
}

func (ur *AdminRepository) Insert(insertAdmin *entity.Admin) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	admin := models.Admin{
		AdminID:  insertAdmin.ID.String(),
		Name:     insertAdmin.Name,
		BelongTo: insertAdmin.BelongStore.ID.String(),
		IsActive: insertAdmin.IsActive,
	}
	err = admin.Insert(ctx, ur.db, boil.Infer())
	if err != nil {
		tran.Rollback()
		return err
	}
	userManager := models.UserManager{
		ID:      insertAdmin.ID.String(),
		IsAdmin: true,
	}

	_, err = userManager.Update(ctx, ur.db, boil.Whitelist(models.UserManagerColumns.IsAdmin))
	if err != nil {
		tran.Rollback()
		return err
	}

	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}
	return nil
}

func (ur *AdminRepository) Update(updateAdmin *entity.Admin) error {
	tran := NewTransaction()
	ctx := context.Background()

	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	admin := models.Admin{
		AdminID:  updateAdmin.ID.String(),
		Name:     updateAdmin.Name,
		BelongTo: updateAdmin.BelongStore.ID.String(),
		IsActive: updateAdmin.IsActive,
	}
	_, err = admin.Update(ctx, ur.db, boil.Infer())
	if err != nil {
		tran.Rollback()
		return err
	}

	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}

func (ur *AdminRepository) Delete(adminID uuid.UUID) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	deleteAdminManager, err := models.FindUserManager(ctx, ur.db, adminID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteAdminManager.Delete(ctx, ur.db)
	if err != nil {
		tran.Rollback()
		return err
	}

	deleteAdminData, err := models.FindAdmin(ctx, ur.db, adminID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteAdminData.Delete(ctx, ur.db)
	if err != nil {
		tran.Rollback()
		return err
	}

	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}
