package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
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

func (pq *CheckinQueryService) GetMyActiveCheckin(userID uuid.UUID) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(userID.String()), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), models.CheckinWhere.Archive.EQ(false), models.CheckinWhere.UserID.EQ(userID.String())).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if checkins == nil {
		return nil, nil
	}
	var result []*entity.Checkin

	for _, coupon := range checkins {
		result = append(result, CheckinModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func (pq *CheckinQueryService) GetMyLastStoreCheckin(userID uuid.UUID, storeID uuid.UUID) (*entity.Checkin, error) {
	checkin, err := models.Checkins(models.CheckinWhere.UserID.EQ(userID.String()), models.CheckinWhere.StoreID.EQ(storeID.String()), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.OrderBy(`checkin_at desc`)).One(context.Background(), pq.db)
	if checkin == nil {
		return nil, nil
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	result := CheckinModelToEntity(checkin, nil, nil)
	return result, nil
}

func (pq *CheckinQueryService) GetMyAllCheckin(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, *types.PageResponse, error) {
	checkins, err := models.Checkins(models.CheckinWhere.UserID.EQ(userID.String()), qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	if checkins == nil {
		return nil, nil, nil
	}

	count, err := models.Checkins(models.CheckinWhere.UserID.EQ(userID.String())).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var result []*entity.Checkin
	for _, checkin := range checkins {
		usr := UserModelToEntity(checkin.R.User, "")
		store := StoreModelToEntity(checkin.R.Store, nil)
		result = append(result, CheckinModelToEntity(checkin, usr, store))
	}
	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}
	return result, pageResponse, nil
}

func (pq *CheckinQueryService) GetAllUserAllCheckin(pager *types.PageQuery) ([]*entity.Checkin, error) {
	checkins, err := models.Checkins(qm.Load(models.CheckinRels.User), qm.Load(models.CheckinRels.Store), qm.Limit(pager.Limit()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if checkins == nil {
		return nil, nil
	}
	var result []*entity.Checkin
	for _, checkin := range checkins {
		usr := UserModelToEntity(checkin.R.User, "")
		store := StoreModelToEntity(checkin.R.Store, nil)
		result = append(result, CheckinModelToEntity(checkin, usr, store))
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
