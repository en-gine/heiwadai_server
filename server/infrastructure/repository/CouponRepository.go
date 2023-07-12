package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ repository.ICouponRepository = &CouponRepository{}

type CouponRepository struct {
	db *sql.DB
}

func NewCouponRepository() (*CouponRepository, error) {
	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	return &CouponRepository{
		db: db,
	}, nil
}

func (pr *CouponRepository) Save(updateCoupon *entity.Coupon) error {
	//user := model.FindUser(context.Background(), pr.db, updateCoupon.UserID.String())

	coupon := models.Coupon{
		ID:                updateCoupon.ID.String(),
		Name:              updateCoupon.Name,
		CouponType:        int(updateCoupon.CouponType),
		DiscountAmount:    int(updateCoupon.DiscountAmount),
		ExpireAt:          updateCoupon.ExpireAt,
		IsCombinationable: updateCoupon.IsCombinationable,
		UsedAt:            null.TimeFromPtr(updateCoupon.UsedAt),
		CreateAt:          updateCoupon.CreateAt,
		CouponStatus:      int(updateCoupon.Status),
	}

	tx, err := pr.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteNotices, err := models.FindCoupon(context.Background(), pr.db, updateCoupon.ID.String())
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = deleteNotices.Delete(context.Background(), pr.db)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, notice := range updateCoupon.Notices {
		modelNotice := models.CouponNotice{
			CouponID: updateCoupon.ID.String(),
			Notice:   notice,
		}
		err = modelNotice.Insert(context.Background(), pr.db, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	deleteStores, err := models.FindCoupon(context.Background(), pr.db, updateCoupon.ID.String())

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = deleteStores.Delete(context.Background(), pr.db)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, store := range updateCoupon.TargetStore {
		modelStore := models.CouponStore{
			CouponID: updateCoupon.ID.String(),
			StoreID:  store.ID.String(),
		}
		err = modelStore.Insert(context.Background(), pr.db, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = coupon.Upsert(context.Background(), pr.db, true, []string{"id"}, boil.Infer(), boil.Infer())

	tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return err
}
