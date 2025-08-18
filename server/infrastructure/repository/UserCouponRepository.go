package repository

import (
	"database/sql"
	"strconv"

	"server/core/entity"
	"server/core/infra/repository"
	"server/db/models"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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

func (pr *UserCouponRepository) Save(tx repository.ITransaction, userCoupon *entity.UserAttachedCoupon) error {
	coupon := models.CouponAttachedUser{
		UserID:   userCoupon.UserID.String(),
		CouponID: userCoupon.Coupon.ID.String(),
		UsedAt:   null.TimeFromPtr(userCoupon.UsedAt),
	}

	err := coupon.Upsert(*tx.Context(), tx.Tran(), true, []string{"coupon_id", "user_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (pr *UserCouponRepository) IssueAll(tx repository.ITransaction, coupon *entity.Coupon, birthMonth *int) (int, error) {
	var sql string = "INSERT INTO " + models.TableNames.CouponAttachedUser +
		" (" + models.CouponAttachedUserColumns.CouponID + ", " +
		models.CouponAttachedUserColumns.UserID + ")" +
		" SELECT '" + coupon.ID.String() + "', " + models.UserDatumColumns.UserID + " FROM " + models.TableNames.UserData

	if birthMonth != nil {
		sql = sql + " WHERE EXTRACT(MONTH FROM " + models.UserDatumColumns.BirthDate + ") = " + strconv.Itoa(*birthMonth)
	}

	res, err := models.NewQuery(
		qm.SQL(sql),
	).ExecContext(*tx.Context(), tx.Tran())
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
