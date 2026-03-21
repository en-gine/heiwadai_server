package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
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

	// First, ensure user_manager record exists (handle race condition with trigger)
	userManager := models.UserManager{
		ID:      updateUser.ID.String(),
		Email:   updateUser.Mail,
		IsAdmin: false,
	}
	err = userManager.Upsert(ctx, tran.Tran(), true, []string{"id"}, boil.Whitelist("email", "update_at"), boil.Infer())
	if err != nil {
		tran.Rollback()
		return fmt.Errorf("failed to upsert user_manager: %w", err)
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
	err = user.Upsert(ctx, tran.Tran(), true, []string{"user_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tran.Rollback()
		return err
	}
	var userOption *models.UserOption
	if updateUserOption != nil {
		userOption = &models.UserOption{
			UserID:          updateUser.ID.String(),
			InnerNote:       updateUserOption.InnerNote,
			IsBlackCustomer: updateUserOption.IsBlackCustomer,
		}
	} else {
		userOption = &models.UserOption{
			UserID:          updateUser.ID.String(),
			InnerNote:       "",
			IsBlackCustomer: false,
		}
	}
	err = userOption.Upsert(ctx, tran.Tran(), true, []string{"user_id"}, boil.Infer(), boil.Infer())
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

func (ur *UserRepository) Delete(userID uuid.UUID) error {
	tran := NewTransaction()
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}

	// CASCADE なしのテーブルを先に削除（user_book, user_report）
	_, err = models.UserBooks(models.UserBookWhere.BookUserID.EQ(userID.String())).DeleteAll(ctx, tran.Tran())
	if err != nil {
		tran.Rollback()
		return err
	}
	_, err = models.UserReports(models.UserReportWhere.UserID.EQ(userID.String())).DeleteAll(ctx, tran.Tran())
	if err != nil {
		tran.Rollback()
		return err
	}

	// user_manager を削除（CASCADE で user_data, user_option, checkin 等も削除される）
	deleteUserManager, err := models.FindUserManager(ctx, tran.Tran(), userID.String())
	if err != nil {
		if err != sql.ErrNoRows {
			tran.Rollback()
			return err
		}
	} else {
		_, err = deleteUserManager.Delete(ctx, tran.Tran())
		if err != nil {
			tran.Rollback()
			return err
		}
	}

	_, err = tran.Exec(fmt.Sprintf("DELETE FROM auth.users WHERE id = '%s'", userID.String()))
	if err != nil {
		if err == sql.ErrNoRows {
			tran.Rollback()
			return nil
		}
		tran.Rollback()
		logger.Error(err.Error())
		return err
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}

func (ur *UserRepository) HasFutureBooking(userID uuid.UUID) (bool, error) {
	ctx := context.Background()
	now := time.Now()
	exists, err := models.UserBooks(
		models.UserBookWhere.BookUserID.EQ(userID.String()),
		models.UserBookWhere.StayTo.GTE(now),
	).Exists(ctx, ur.db)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *UserRepository) DeleteUnderRegisterUser(userID uuid.UUID) error {
	_, err := ur.db.Exec(fmt.Sprintf("DELETE FROM auth.users WHERE id = '%s'", userID.String()))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error(err.Error())
		return err
	}
	return nil
}
