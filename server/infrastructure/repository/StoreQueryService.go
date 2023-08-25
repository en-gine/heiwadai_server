package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/db/models"

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
	store, err := models.FindStore(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}

	return StoreModelToEntity(store), nil
}

func (pq *StoreQueryService) GetActiveAll() ([]*entity.Store, error) {
	stores, err := models.Stores(models.StoreWhere.IsActive.EQ(true)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store))
	}
	return result, nil
}

func (pq *StoreQueryService) GetAll() ([]*entity.Store, error) {
	stores, err := models.Stores().All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Store
	for _, store := range stores {
		result = append(result, StoreModelToEntity(store))
	}
	return result, nil

}

func StoreModelToEntity(model *models.Store) *entity.Store {
	infoModel, _ := model.StayableStoreInfo().One(context.Background(), InitDB())
	stayableStoreInfo := entity.RegenStayableStoreInfo(
		infoModel.Parking,
		infoModel.Latitude,
		infoModel.Longitude,
		infoModel.AccessInfo,
		infoModel.RestAPIURL,
		infoModel.BookingSystemID,
	)
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
		stayableStoreInfo,
		model.IsActive,
		uuid.MustParse(model.QRCode),
		uuid.MustParse(model.UnLimitedQRCode),
	)
}
