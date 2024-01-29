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
	ctx := context.Background()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}
	tran.Tx.QueryContext(ctx, "SET CONSTRAINTS ALL DEFERRED;") // トランザクション内で外部キー制約を無効化
	err = store.Upsert(ctx, tran.Tx, true, []string{"id"}, boil.Infer(), boil.Infer())
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
		err = StayableStoreInfo.Upsert(ctx, tran.Tx, true, []string{"store_id"}, boil.Infer(), boil.Infer())
		if err != nil {
			tran.Rollback()
			return err
		}
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}

func (pr *StoreRepository) Delete(storeID uuid.UUID) error {
	ctx := context.Background()
	tran := NewTransaction()
	err := tran.Begin(ctx)
	if err != nil {
		return err
	}

	deleteStore, err := models.FindStore(ctx, tran.Tx, storeID.String())
	if err != nil {
		return err
	}
	_, err = deleteStore.Delete(ctx, tran.Tx)
	if err != nil {
		tran.Rollback()
		return err
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return err
	}

	return nil
}

func (pr *StoreRepository) RegenQR(storeID uuid.UUID) (*uuid.UUID, error) {
	ctx := context.Background()
	tran := NewTransaction()
	err := tran.Begin(ctx)
	if err != nil {
		return nil, err
	}

	store, err := models.FindStore(ctx, tran.Tx, storeID.String())
	if err != nil {
		return nil, err
	}
	qrcode, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	store.QRCode = qrcode.String()
	_, err = store.Update(ctx, tran.Tx, boil.Infer())
	if err != nil {
		tran.Rollback()
		return nil, err
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return nil, err
	}

	return &qrcode, nil
}

func (pr *StoreRepository) RegenUnlimitQR(storeID uuid.UUID) (*uuid.UUID, error) {
	ctx := context.Background()
	tran := NewTransaction()
	err := tran.Begin(ctx)
	if err != nil {
		return nil, err
	}

	store, err := models.FindStore(ctx, tran.Tx, storeID.String())
	if err != nil {
		return nil, err
	}
	qrcode, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	store.UnLimitedQRCode = qrcode.String()
	_, err = store.Update(ctx, tran.Tx, boil.Infer())
	if err != nil {
		tran.Rollback()
		return nil, err
	}
	err = tran.Commit()
	if err != nil {
		tran.Rollback()
		return nil, err
	}

	return &qrcode, nil
}
