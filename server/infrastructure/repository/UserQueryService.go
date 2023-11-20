package repository

import (
	"context"
	"database/sql"
	"time"

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

func (pq *UserQueryService) GetByID(id uuid.UUID) (*entity.User, error) {
	user, err := models.FindUserDatum(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return UserModelToEntity(user, user.R.User.Email), nil
}

func (pq *UserQueryService) GetOptionByID(id uuid.UUID) (*entity.UserOption, error) {
	option, err := models.FindUserOption(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &entity.UserOption{
		InnerNote:       option.InnerNote,
		IsBlackCustomer: option.IsBlackCustomer,
	}, nil
}

func (pq *UserQueryService) GetByMail(mail string) (*entity.User, error) {
	usermanager, err := models.UserManagers(models.UserManagerWhere.Email.EQ(mail)).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return UserModelToEntity(usermanager.R.UserUserDatum, usermanager.Email), nil
}

func (pq *UserQueryService) GetMailOKUser(prefectures *[]entity.Prefecture) ([]*entity.User, error) {
	var users []*models.UserDatum
	var err error

	qm := GetMailUserWhereMods(prefectures)
	users, err = models.UserData(qm...).All(context.Background(), pq.db)

	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	var entities []*entity.User
	for _, user := range users {
		entity := UserModelToEntity(user, user.R.User.Email)
		entities = append(entities, entity)
	}
	return entities, nil
}

func (pq *UserQueryService) GetMailOKUserCount(prefectures *[]entity.Prefecture) (*int, error) {
	var err error
	var count int
	qm := GetMailUserWhereMods(prefectures)
	int64Count, err := models.UserData(qm...).Count(context.Background(), pq.db)
	count = int(int64Count)

	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		count = 0
		return &count, nil
	}
	return &count, nil
}

// CountとAll両方で使えるようにクエリのみ返す
func GetMailUserWhereMods(prefectures *[]entity.Prefecture) []qm.QueryMod {
	var preIds []int

	if prefectures == nil {
		return []qm.QueryMod{models.UserDatumWhere.AcceptMail.EQ(true)}
	} else {
		for _, prefecture := range *prefectures {
			preIds = append(preIds, prefecture.ToInt())
		}
		return []qm.QueryMod{models.UserDatumWhere.Prefecture.IN(preIds), models.UserDatumWhere.AcceptMail.EQ(true)}
	}
}

func (pq *UserQueryService) GetList(query *types.UserQuery, pager *types.PageQuery) ([]*entity.UserWichLastCheckin, *types.PageResponse, error) {
	var firstNameQuery qm.QueryMod = nil
	if query.FirstName != nil {
		firstNameQuery = models.UserDatumWhere.FirstName.EQ("%" + *query.FirstName + "%")
	}
	var lastNameQuery qm.QueryMod = nil
	if query.FirstName != nil {
		lastNameQuery = models.UserDatumWhere.LastName.EQ("%" + *query.LastName + "%")
	}
	var firstNameKanaQuery qm.QueryMod = nil
	if query.FirstNameKana != nil {
		firstNameKanaQuery = models.UserDatumWhere.FirstNameKana.EQ("%" + *query.FirstNameKana + "%")
	}
	var lastNameKanaQuery qm.QueryMod = nil
	if query.LastNameKana != nil {
		lastNameKanaQuery = models.UserDatumWhere.LastNameKana.EQ("%" + *query.LastNameKana + "%")
	}
	var prefectureQuery qm.QueryMod = nil
	if query.LastNameKana != nil {
		prefectureQuery = models.UserDatumWhere.Prefecture.EQ(query.Prefecture.ToInt())
	}

	userdata, err := models.UserData(
		firstNameQuery, lastNameQuery, firstNameKanaQuery, lastNameKanaQuery, prefectureQuery,
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset()),
		qm.Load(models.UserDatumRels.User),
		qm.Load(models.UserDatumRels.UserCheckins, qm.OrderBy(models.CheckinColumns.CheckInAt+" desc"), qm.Limit(1)),
	).All(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil, nil
	}
	var result []*entity.UserWichLastCheckin
	for _, user := range userdata {
		u := UserModelToEntity(user, user.R.User.Email)
		var lastCheckinAt *time.Time = nil
		lastCheckIn := user.R.UserCheckins[0]
		if lastCheckIn != nil {
			lastCheckinAt = &lastCheckIn.CheckInAt
		}
		userWithCheckin := &entity.UserWichLastCheckin{
			User:          u,
			LastCheckinAt: lastCheckinAt,
		}

		result = append(result, userWithCheckin)
	}

	count, err := models.UserData(firstNameQuery, lastNameQuery, firstNameKanaQuery, lastNameKanaQuery, prefectureQuery).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil, nil
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = &types.PageResponse{
			CurrentPage: *pager.CurrentPage,
			PerPage:     *pager.PerPage,
			TotalCount:  int(count),
			TotalPage:   int(count) / *pager.PerPage,
		}
	}

	return result, pageResponse, nil
}

func UserModelToEntity(model *models.UserDatum, email string) *entity.User {
	return entity.RegenUser(
		uuid.MustParse(model.UserID),
		model.FirstName,
		model.LastName,
		model.FirstNameKana,
		model.LastNameKana,
		model.CompanyName.Ptr(),
		model.BirthDate.Ptr(),
		&model.ZipCode.String,
		entity.Prefecture(model.Prefecture),
		model.City.Ptr(),
		model.Address.Ptr(),
		model.Tel.Ptr(),
		email,
		model.AcceptMail,
	)
}
