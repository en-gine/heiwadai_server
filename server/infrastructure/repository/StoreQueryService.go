package repository

import (
	"context"
	"database/sql"
	"errors"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/google/uuid"
)

var _ queryservice.IStoreQueryService = &StoreQueryService{}

type StoreQueryService struct {
	db *sql.DB
}

func NewStoreQueryService() *StoreQueryService {
	db := InitDB()

	return &StoreQueryService{
		db: db,
	}
}

func (pq *StoreQueryService) GetByID(id uuid.UUID) (*entity.Store, error) {
	ctx := context.Background()
	store, err := models.FindStore(ctx, pq.db, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if store == nil {
		return nil, nil
	}
	info, err := models.StayableStoreInfos(models.StayableStoreInfoWhere.StoreID.EQ(store.ID)).One(ctx, pq.db)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(err.Error())
		return nil, err
	}
	if info == nil {
		info = nil
	}

	return StoreModelToEntity(store, info), nil
}

func (pq *StoreQueryService) GetActiveAll() ([]*entity.Store, error) {
	stores, err := models.Stores(qm.Load(models.StoreRels.StayableStoreInfo), models.StoreWhere.IsActive.EQ(true), qm.OrderBy(models.StoreColumns.CreateAt+" ASC")).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store, store.R.StayableStoreInfo))
	}
	return result, nil
}

func (pq *StoreQueryService) GetStayableByID(id uuid.UUID) (*entity.StayableStore, error) {
	store, err := models.FindStore(context.Background(), pq.db, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if store == nil {
		return nil, nil
	}
	infoModel, err := store.StayableStoreInfo().One(context.Background(), InitDB())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	return StayableStoreToEntity(store, infoModel), nil
}

func (pq *StoreQueryService) GetStayables() ([]*entity.StayableStore, error) {
	stores, err := models.Stores(models.StoreWhere.IsActive.EQ(true), models.StoreWhere.Stayable.EQ(true)).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.StayableStore
	for _, store := range stores {
		infoModel, err := store.StayableStoreInfo().One(context.Background(), pq.db)
		if err != nil || infoModel == nil {
			// StayableStoreInfoがないストアはスキップ
			continue
		}
		stayable := StayableStoreToEntity(store, infoModel)
		result = append(result, stayable)
	}
	return result, nil
}

func (pq *StoreQueryService) GetAll() ([]*entity.Store, error) {
	stores, err := models.Stores(qm.Load(models.StoreRels.StayableStoreInfo)).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if stores == nil {
		return nil, nil
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store, store.R.StayableStoreInfo))
	}
	return result, nil
}

func (pq *StoreQueryService) GetStayableByBookingID(bookingID string) (*entity.StayableStore, error) {
	stayables, err := pq.GetStayables()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	for _, stayable := range stayables {
		if stayable.BookingSystemID == bookingID {
			return stayable, nil
		}
	}
	return nil, errors.New("該当のStayableStoreがBookingIDから見つけることが出来ません。")
}

func (pq *StoreQueryService) GetStoreByQrCode(hash uuid.UUID) (*entity.Store, error) {
	store, err := models.Stores(models.StoreWhere.QRCode.EQ(hash.String())).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if store == nil {
		return nil, nil
	}

	return StoreModelToEntity(store, nil), nil
}

func (pq *StoreQueryService) GetStoreByUnlimitQrCode(hash uuid.UUID) (*entity.Store, error) {
	store, err := models.Stores(models.StoreWhere.UnLimitedQRCode.EQ(hash.String())).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if store == nil {
		return nil, nil
	}

	return StoreModelToEntity(store, nil), nil
}

func StoreModelToEntity(model *models.Store, info *models.StayableStoreInfo) *entity.Store {
	var stayableInfo *entity.StayableStoreInfo
	if info != nil {
		stayableInfo = StayableInfoToEntity(info)
	}

	return entity.RegenStore(
		uuid.MustParse(model.ID),
		model.Name,
		model.BranchName.Ptr(),
		model.ZipCode,
		model.Address,
		model.Tel,
		model.SiteURL,
		model.StampImageURL,
		model.Stayable,
		model.IsActive,
		uuid.MustParse(model.QRCode),
		uuid.MustParse(model.UnLimitedQRCode),
		stayableInfo,
	)
}

func StayableInfoToEntity(info *models.StayableStoreInfo) *entity.StayableStoreInfo {
	return entity.RegenStayableStoreInfo(
		info.Parking,
		info.Latitude,
		info.Longitude,
		info.AccessInfo,
		info.RestAPIURL,
		info.BookingSystemID,
		info.BookingSystemLoginID.String,
		info.BookingSystemPassword.String,
	)
}

func StayableStoreToEntity(store *models.Store, info *models.StayableStoreInfo) *entity.StayableStore {
	var stayableInfo *entity.StayableStoreInfo
	if info != nil {
		stayableInfo = StayableInfoToEntity(info)
	}
	var storeEntity *entity.Store
	if store != nil {
		storeEntity = StoreModelToEntity(store, nil)
	}

	return entity.RegenStayableStore(
		storeEntity,
		stayableInfo,
	)
}
