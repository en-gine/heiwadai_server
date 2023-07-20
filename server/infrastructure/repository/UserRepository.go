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

func (pr *UserRepository) Save(updateUser *entity.User) error {
	post := models.User{
		ID:            updateUser.ID.String(),
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
		Mail:          updateUser.Mail,
		AcceptMail:    updateUser.AcceptMail,
	}
	err := post.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	return err
}

func (pr *UserRepository) Delete(postId uuid.UUID) error {
	deleteUser, err := models.FindUser(context.Background(), pr.db, postId.String())
	if err != nil {
		return err
	}
	_, err = deleteUser.Delete(context.Background(), pr.db)
	return err
}
