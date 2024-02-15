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

func NewCouponRepository() *CouponRepository {
	db := InitDB()

	return &CouponRepository{
		db: db,
	}
}

func (pr *CouponRepository) Save(updateCoupon *entity.Coupon) error {
	var count int
	if updateCoupon == nil {
		count = 0
	} else {
		count = *updateCoupon.IssueCount
	}
	coupon := models.Coupon{
		ID:                updateCoupon.ID.String(),
		Name:              updateCoupon.Name,
		CouponType:        int(updateCoupon.CouponType),
		DiscountAmount:    int(updateCoupon.DiscountAmount),
		Notices:           updateCoupon.Notices,
		ExpireAt:          updateCoupon.ExpireAt,
		IsCombinationable: updateCoupon.IsCombinationable,
		CreateAt:          updateCoupon.CreateAt,
		CouponStatus:      int(updateCoupon.Status),
		IssueCount:        count,
		IssueAt:           null.TimeFromPtr(updateCoupon.IssueAt),
	}

	ctx := context.Background()
	tx := NewTransaction()
	err := tx.Begin(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// _, err = tx.Tx.ExecContext(ctx, "SET CONSTRAINTS ALL DEFERRED;") // トランザクション内で外部キー制約を無効化
	if err != nil {
		tx.Rollback()
		return err
	}

	err = coupon.Upsert(ctx, tx.Tx, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tx.Rollback()
		return err
	}

	var modelStores []*models.Store
	for _, store := range updateCoupon.TargetStore {
		modelStore := &models.Store{
			ID: store.ID.String(),
		}
		modelStores = append(modelStores, modelStore)
	}
	err = coupon.SetStores(ctx, tx.Tx, false, modelStores...)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
