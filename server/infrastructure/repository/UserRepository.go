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
	tran.Begin()
	user := models.UserDatum{
		UserID:        updateUser.ID.String(),
		FirstName:     updateUser.FirstName,
		LastName:      updateUser.LastName,
		FirstNameKana: updateUser.FirstNameKana,
		LastNameKana:  updateUser.LastNameKana,
		CompanyName:   null.StringFromPtr(updateUser.CompanyName),
		BirthDate:     updateUser.BirthDate,
		ZipCode:       null.StringFromPtr(updateUser.ZipCode),
		Prefecture:    int(updateUser.Prefecture),
		City:          null.StringFromPtr(updateUser.City),
		Address:       null.StringFromPtr(updateUser.Address),
		Tel:           null.StringFromPtr(updateUser.Tel),
		AcceptMail:    updateUser.AcceptMail,
	}
	err := user.Upsert(context.Background(), ur.db, true, []string{"user_id"}, boil.Infer(), boil.Infer())
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
		err := userOption.Upsert(context.Background(), ur.db, true, []string{"user_id"}, boil.Infer(), boil.Infer())
		if err != nil {
			tran.Rollback()
			return err
		}
	}
	tran.Commit()
	return nil
}

func (ur *UserRepository) Delete(userID uuid.UUID) error {

	tran := NewTransaction()
	tran.Begin()
	deleteUserManager, err := models.FindUserManager(context.Background(), ur.db, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserManager.Delete(context.Background(), ur.db)
	if err != nil {
		tran.Rollback()
		return err
	}

	deleteUserData, err := models.FindUserDatum(context.Background(), ur.db, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserData.Delete(context.Background(), ur.db)
	if err != nil {
		tran.Rollback()
		return err
	}
	deleteUserOption, err := models.FindUserOption(context.Background(), ur.db, userID.String())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = deleteUserOption.Delete(context.Background(), ur.db)
	if err != nil {
		tran.Rollback()
		return err
	}
	tran.Commit()
	return nil
}
