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

func NewStoreQueryService() (*StoreQueryService, error) {
	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	return &StoreQueryService{
		db: db,
	}, nil
}

func (pq *StoreQueryService) GetById(id uuid.UUID) (*entity.Store, error) {
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
	return entity.RegenStore(
		uuid.MustParse(model.ID),
		model.Name,
		model.Address,
		model.Tel,
		model.Parking,
		model.AccessInfo,
		model.IsActive,
		uuid.MustParse(model.QRCode),
		uuid.MustParse(model.UnLimitedQRCode),
	)
}
