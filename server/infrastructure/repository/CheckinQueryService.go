package repository

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

func NewCheckinQueryService() *CheckinQueryService {
	db := InitDB()

	return &CheckinQueryService{
		db: db,
	}
}

func (pq *CheckinQueryService) GetActiveCheckin(userID uuid.UUID) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(userID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), models.CheckinWhere.Archive.EQ(false), models.CheckinWhere.UserID.EQ(null.StringFrom(userID.String()))).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	var result []*entity.Checkin

	for _, coupon := range checkins {
		result = append(result, CheckinModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func (pq *CheckinQueryService) GetLastStoreCheckin(userID uuid.UUID, storeID uuid.UUID) (*entity.Checkin, error) {
	checkin, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(userID.String())), models.CheckinWhere.StoreID.EQ(null.StringFrom(storeID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.OrderBy(`checkin_at desc`)).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	result := CheckinModelToEntity(checkin, nil, nil)
	return result, nil
}

func (pq *CheckinQueryService) GetAllCheckin(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(null.StringFrom(userID.String())), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
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
