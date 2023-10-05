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

var _ repository.IUserCouponRepository = &UserCouponRepository{}

type UserCouponRepository struct {
	db *sql.DB
}

func NewUserCouponRepository() *UserCouponRepository {
	db := InitDB()

	return &UserCouponRepository{
		db: db,
	}
}

func (pr *UserCouponRepository) Save(ctx context.Context, userCoupon *entity.UserAttachedCoupon) error {
	coupon := models.CouponAttachedUser{
		UserID:   userCoupon.UserID.String(),
		CouponID: userCoupon.Coupon.ID.String(),
		UsedAt:   null.TimeFromPtr(userCoupon.UsedAt),
	}
	tx := NewTransaction()
	err := tx.Begin(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = coupon.Upsert(ctx, pr.db, true, []string{"coupon_id", "user_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}
