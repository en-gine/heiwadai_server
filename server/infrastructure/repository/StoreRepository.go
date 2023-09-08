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

var _ repository.IStoreRepository = &StoreRepository{}

type StoreRepository struct {
	db *sql.DB
}

func NewStoreRepository() *StoreRepository {
	db := InitDB()

	return &StoreRepository{
		db: db,
	}
}

func (pr *StoreRepository) Save(updateStore *entity.Store, stayableInfo *entity.StayableStoreInfo) error {

	store := models.Store{
		ID:              updateStore.ID.String(),
		Name:            updateStore.Name,
		BranchName:      null.StringFromPtr(updateStore.BranchName),
		ZipCode:         updateStore.ZipCode,
		Address:         updateStore.Address,
		Tel:             updateStore.Tel,
		SiteURL:         updateStore.SiteURL,
		StampImageURL:   updateStore.StampImageURL,
		Stayable:        updateStore.Stayable,
		IsActive:        updateStore.IsActive,
		QRCode:          updateStore.QRCode.String(),
		UnLimitedQRCode: updateStore.UnLimitedQRCode.String(),
	}

	tran := NewTransaction()
	tran.Begin()
	err := store.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tran.Rollback()
		return err
	}

	if stayableInfo != nil {
		StayableStoreInfo := &models.StayableStoreInfo{
			StoreID:         updateStore.ID.String(),
			Parking:         stayableInfo.Parking,
			Latitude:        stayableInfo.Latitude,
			Longitude:       stayableInfo.Longitude,
			AccessInfo:      stayableInfo.AccessInfo,
			RestAPIURL:      stayableInfo.RestAPIURL,
			BookingSystemID: stayableInfo.BookingSystemID,
		}
		err = StayableStoreInfo.Upsert(context.Background(), pr.db, true, []string{"store_id"}, boil.Infer(), boil.Infer())
		if err != nil {
			tran.Rollback()
			return err
		}
	}
	tran.Commit()
	return err
}

func (pr *StoreRepository) Delete(storeID uuid.UUID) error {
	deleteStore, err := models.FindStore(context.Background(), pr.db, storeID.String())
	if err != nil {
		return err
	}
	tran := NewTransaction()
	tran.Begin()
	_, err = deleteStore.Delete(context.Background(), pr.db)
	if err != nil {
		tran.Rollback()
		return err
	}
	tran.Commit()
	return err
}
