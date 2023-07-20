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

var _ queryservice.IUserQueryService = &UserQueryService{}

type UserQueryService struct {
	db *sql.DB
}

func NewUserQueryService() *UserQueryService {
	db := InitDB()

	return &UserQueryService{
		db: db,
	}
}

func (pq *UserQueryService) GetById(id uuid.UUID) (*entity.User, error) {
	user, err := models.FindUser(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}

	return UserModelToEntity(user), nil
}
func (pq *UserQueryService) GetByMail(mail string) (*entity.User, error) {
	user, err := models.Users(models.UserWhere.Mail.EQ(mail)).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return UserModelToEntity(user), nil
}

func (pq *UserQueryService) GetUserByPrefecture(prefectures []*entity.Prefecture) ([]*entity.User, error) {
	var preIds []int
	for _, prefecture := range prefectures {
		preIds = append(preIds, prefecture.ToInt())
	}

	users, err := models.Users(models.UserWhere.Prefecture.IN(preIds)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var entities []*entity.User
	for _, u := range users {
		entity := UserModelToEntity(u)
		entities = append(entities, entity)
	}
	return entities, nil

}

func (pq *UserQueryService) GetAll(pager *types.PageQuery) ([]*entity.User, error) {
	users, err := models.Users(qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	for _, user := range users {
		result = append(result, UserModelToEntity(user))
	}
	return result, nil

}

func UserModelToEntity(model *models.User) *entity.User {
	return entity.RegenUser(
		uuid.MustParse(model.ID),
		model.FirstName,
		model.LastName,
		model.FirstNameKana,
		model.LastNameKana,
		model.CompanyName.Ptr(),
		model.BirthDate,
		&model.ZipCode.String,
		entity.Prefecture(model.Prefecture),
		model.City.Ptr(),
		model.Address.Ptr(),
		model.Tel.Ptr(),
		model.Mail,
		model.AcceptMail,
	)
}
