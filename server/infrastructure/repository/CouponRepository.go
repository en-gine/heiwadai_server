package repository

import (
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

func (pr *CouponRepository) Save(tx repository.ITransaction, updateCoupon *entity.Coupon) error {
	var count int
	if updateCoupon == nil || updateCoupon.IssueCount == nil {
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

	err := coupon.Upsert(*tx.Context(), tx.Tran(), true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	var modelStores []*models.Store
	for _, store := range updateCoupon.TargetStore {
		modelStore := &models.Store{
			ID: store.ID.String(),
		}
		modelStores = append(modelStores, modelStore)
	}
	err = coupon.SetStores(*tx.Context(), tx.Tran(), false, modelStores...)
	if err != nil {
		return err
	}

	return nil
}
