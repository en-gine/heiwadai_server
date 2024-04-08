package repository

import (
	"context"
	"database/sql"
	"fmt"

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
	defer func() {
		if err != nil {
			tran.Rollback()
		} else {
			err = tran.Commit()
		}
	}()
	admin := models.Admin{
		AdminID:  insertAdmin.ID.String(),
		Name:     insertAdmin.Name,
		BelongTo: insertAdmin.BelongStore.ID.String(),
		IsActive: insertAdmin.IsActive,
	}
	err = admin.Insert(ctx, tran.Tran(), boil.Infer())
	if err != nil {
		return err
	}
	userManager := models.UserManager{
		ID:      insertAdmin.ID.String(),
		IsAdmin: true,
	}

	_, err = userManager.Update(ctx, tran.Tran(), boil.Whitelist(models.UserManagerColumns.IsAdmin))
	if err != nil {
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
	defer func() {
		if err != nil {
			tran.Rollback()
		} else {
			err = tran.Commit()
		}
	}()

	admin := models.Admin{
		AdminID:  updateAdmin.ID.String(),
		Name:     updateAdmin.Name,
		BelongTo: updateAdmin.BelongStore.ID.String(),
		IsActive: updateAdmin.IsActive,
	}
	_, err = admin.Update(ctx, tran.Tran(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (ur *AdminRepository) Delete(adminID uuid.UUID) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	defer func() {
		if err != nil {
			tran.Rollback()
		} else {
			err = tran.Commit()
		}
	}()
	if err != nil {
		return err
	}
	// deleteAdminManager, err := models.FindUserManager(ctx, tran.Tran(), adminID.String())
	// if err != nil {
	// 	return err
	// }
	// _, err = deleteAdminManager.Delete(ctx, tran.Tran())
	// if err != nil {
	// 	return err
	// }
	_, err = models.FindAdmin(ctx, tran.Tran(), adminID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	// _, err = deleteAdminData.Delete(ctx, tran.Tran())
	// if err != nil {
	// 	return err
	// }
	_, err = tran.Exec(fmt.Sprintf("DELETE FROM auth.users WHERE id = '%s'", adminID.String()))
	if err != nil {
		return err
	}

	return nil
}
