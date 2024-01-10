package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.IUserRepository = &UserRepository{}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	db := InitDB()

	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(updateUser *entity.User, updateUserOption *entity.UserOption) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	user := models.UserDatum{
		UserID:        updateUser.ID.String(),
		FirstName:     updateUser.FirstName,
		LastName:      updateUser.LastName,
		FirstNameKana: updateUser.FirstNameKana,
		LastNameKana:  updateUser.LastNameKana,
		CompanyName:   null.StringFromPtr(updateUser.CompanyName),
		BirthDate:     null.TimeFromPtr(updateUser.BirthDate),
		ZipCode:       null.StringFromPtr(updateUser.ZipCode),
		Prefecture:    int(updateUser.Prefecture),
		City:          null.StringFromPtr(updateUser.City),
		Address:       null.StringFromPtr(updateUser.Address),
		Tel:           null.StringFromPtr(updateUser.Tel),
		AcceptMail:    updateUser.AcceptMail,
	}
	err = user.Upsert(ctx, tran.Tx, true, []string{"user_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tran.Rollback()
		return err
	}
	if updateUserOption != nil {
		userOption := models.UserOption{
			UserID:          updateUser.ID.String(),
			InnerNote:       updateUserOption.InnerNote,
			IsBlackCustomer: updateUserOption.IsBlackCustomer,
		}
		err := userOption.Upsert(ctx, tran.Tx, true, []string{"user_id"}, boil.Infer(), boil.Infer())
		if err != nil {
			tran.Rollback()
			return err
		}
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(userID uuid.UUID) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	deleteUserManager, err := models.FindUserManager(ctx, tran.Tx, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserManager.Delete(ctx, tran.Tx)
	if err != nil {
		tran.Rollback()
		return err
	}

	deleteUserData, err := models.FindUserDatum(ctx, tran.Tx, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserData.Delete(ctx, tran.Tx)
	if err != nil {
		tran.Rollback()
		return err
	}
	deleteUserOption, err := models.FindUserOption(ctx, tran.Tx, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserOption.Delete(ctx, tran.Tx)
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
