package repository

import (
	"context"
	"database/sql"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

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
	// user, err := models.FindUserDatum(context.Background(), pq.db, id.String())
	user, err := models.UserData(
		models.UserDatumWhere.UserID.EQ(id.String()),
		qm.Load(models.UserDatumRels.User),
	).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if user == nil {
		return nil, nil
	}
	return UserModelToEntity(user, user.R.User.Email), nil
}

func (pq *UserQueryService) GetOptionByID(id uuid.UUID) (*entity.UserOption, error) {
	option, err := models.FindUserOption(context.Background(), pq.db, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if option == nil {
		return nil, nil
	}
	return &entity.UserOption{
		InnerNote:       option.InnerNote,
		IsBlackCustomer: option.IsBlackCustomer,
	}, nil
}

func (pq *UserQueryService) GetByMail(mail string) (*entity.User, error) {
	usermanager, err := models.UserManagers(models.UserManagerWhere.Email.EQ(mail), qm.Load(models.UserManagerRels.UserUserDatum), models.UserManagerWhere.IsAdmin.EQ(false)).One(context.Background(), pq.db)

	if usermanager == nil {
		return nil, nil
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	return UserModelToEntity(usermanager.R.UserUserDatum, usermanager.Email), nil
}

func (pq *UserQueryService) GetMailOKUser(prefectures *[]entity.Prefecture) ([]*entity.User, error) {
	var users []*models.UserDatum
	var err error

	queryMods := GetMailUserWhereMods(prefectures)
	queryMods = append(queryMods, qm.Load(models.UserDatumRels.User))
	users, err = models.UserData(queryMods...).All(context.Background(), pq.db)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	if users == nil {
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
	queryMods := GetMailUserWhereMods(prefectures)
	int64Count, err := models.UserData(queryMods...).Count(context.Background(), pq.db)
	count = int(int64Count)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	return &count, nil
}

func GetMailUserWhereMods(prefectures *[]entity.Prefecture) []qm.QueryMod {
	// CountとAll両方で使えるようにクエリのみ返す
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
	userFilterMods := GetUserListFilterMods(query)
	userListMods := []qm.QueryMod{
		qm.Load(models.UserDatumRels.User),
		qm.Load(models.UserDatumRels.UserCheckins,
			qm.OrderBy(models.CheckinColumns.CheckInAt+" desc"),
			qm.Limit(1)),
	}
	userdata, err := models.UserData(
		append(userFilterMods, userListMods...)...,
	).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}
	if userdata == nil {
		return nil, nil, nil
	}
	var result []*entity.UserWichLastCheckin
	for _, user := range userdata {
		u := UserModelToEntity(user, user.R.User.Email)
		var lastCheckIn *models.Checkin = nil
		var lastCheckinAt *time.Time = nil
		if user.R.UserCheckins != nil {
			lastCheckIn = user.R.UserCheckins[0]
		}
		if lastCheckIn != nil {
			lastCheckinAt = &lastCheckIn.CheckInAt
		}
		userWithCheckin := &entity.UserWichLastCheckin{
			User:          u,
			LastCheckinAt: lastCheckinAt,
		}

		result = append(result, userWithCheckin)
	}

	count, err := models.UserData(userFilterMods...).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}

	return result, pageResponse, nil
}

func GetUserListFilterMods(query *types.UserQuery) []qm.QueryMod {
	var qmMods []qm.QueryMod
	// CountとAll両方で使えるようにクエリのみ返す
	var firstNameQuery qm.QueryMod
	var lastNameQuery qm.QueryMod
	var firstNameKanaQuery qm.QueryMod
	var lastNameKanaQuery qm.QueryMod
	var prefectureQuery qm.QueryMod

	if query != nil {
		if query.FirstName != nil {
			firstNameQuery = models.UserDatumWhere.FirstName.EQ("%" + *query.FirstName + "%")
			qmMods = append(qmMods, firstNameQuery)
		}
		if query.LastName != nil {
			lastNameQuery = models.UserDatumWhere.LastName.EQ("%" + *query.LastName + "%")
			qmMods = append(qmMods, lastNameQuery)
		}
		if query.FirstNameKana != nil {
			firstNameKanaQuery = models.UserDatumWhere.FirstNameKana.EQ("%" + *query.FirstNameKana + "%")
			qmMods = append(qmMods, firstNameKanaQuery)
		}
		if query.LastNameKana != nil {
			lastNameKanaQuery = models.UserDatumWhere.LastNameKana.EQ("%" + *query.LastNameKana + "%")
			qmMods = append(qmMods, lastNameKanaQuery)
		}
		if query.Prefecture != nil {
			prefectureQuery = models.UserDatumWhere.Prefecture.EQ(query.Prefecture.ToInt())
			qmMods = append(qmMods, prefectureQuery)
		}
	}

	return qmMods
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
