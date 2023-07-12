package implement

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.ICheckinQueryService = &CheckinQueryService{}

type CheckinQueryService struct {
	db *sql.DB
}

func NewCheckinQueryService() (*CheckinQueryService, error) {
	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	return &CheckinQueryService{
		db: db,
	}, nil
}

func (pq *CheckinQueryService) GetActiveCheckin(user *entity.User) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(user.ID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), models.CheckinWhere.Archive.EQ(false), models.CheckinWhere.UserID.EQ(null.StringFrom(user.ID.String()))).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Checkin

	for _, coupon := range checkins {
		result = append(result, CheckinModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func (pq *CheckinQueryService) GetLastStoreCheckin(user *entity.User, store *entity.Store) (*entity.Checkin, error) {
	checkin, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(user.ID.String())), models.CheckinWhere.StoreID.EQ(null.StringFrom(store.ID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.OrderBy(`checkin_at desc`)).One(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result = CheckinModelToEntity(checkin, nil, nil)
	return result, nil
}

func (pq *CheckinQueryService) GetAllCheckin(user *entity.User, pager *types.PageQuery) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(user.ID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Checkin
	for _, coupon := range checkins {
		result = append(result, CheckinModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func CheckinModelToEntity(model *models.Checkin, user *entity.User, store *entity.Store) *entity.Checkin {
	return entity.RegenCheckin(
		uuid.MustParse(model.ID),
		store,
		user,
		model.CheckInAt,
		model.Archive,
	)
}
