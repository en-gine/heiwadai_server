package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.IUserCouponQueryService = &UserCouponQueryService{}

type UserCouponQueryService struct {
	db *sql.DB
}

func NewUserCouponQueryService() *UserCouponQueryService {
	db := InitDB()
	return &UserCouponQueryService{
		db: db,
	}
}

func (pq *UserCouponQueryService) GetByID(userID uuid.UUID, couponID uuid.UUID) (*entity.UserAttachedCoupon, error) {
	// userCoupon, err := models.FindCouponAttachedUser(context.Background(), pq.db, couponID.String(), userID.String())
	userCoupon, err := models.CouponAttachedUsers(
		models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		models.CouponAttachedUserWhere.CouponID.EQ(couponID.String())).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	if userCoupon == nil {
		return nil, errors.New("該当のクーポンIDが見つかりません。")
	}
	coupon := userCoupon.R.Coupon
	if coupon == nil {
		return nil, errors.New("該当のクーポンが見つかりません。")
	}
	entityCoupon := CouponModelToEntity(coupon, nil)
	return entity.RegenUserAttachedCoupon(
		userID,
		entityCoupon,
		userCoupon.UsedAt.Ptr(),
	), nil
}

func (pq *UserCouponQueryService) GetActiveAll(userID uuid.UUID) ([]*entity.UserAttachedCoupon, error) {

	userCoupons, err := models.CouponAttachedUsers(
		models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		models.CouponAttachedUserWhere.UsedAt.IsNull(),
		qm.Load(models.CouponAttachedUserRels.Coupon), // リレーションをロード
		qm.InnerJoin(models.TableNames.Coupon+" ON "+models.TableNames.CouponAttachedUser+"."+models.CouponAttachedUserColumns.CouponID+" = "+models.TableNames.Coupon+".id"),
		models.CouponWhere.ExpireAt.GT(time.Now().AddDate(0, 0, -1)), // 親テーブルの条件を指定
		qm.OrderBy(models.CouponAttachedUserColumns.CouponID+" DESC"),
	).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if userCoupons == nil {
		return nil, nil
	}
	var result = ModelToUserAttachedCoupon(userID, &userCoupons)
	return result, nil
}

func (pq *UserCouponQueryService) GetExpires(userID uuid.UUID, limit int) ([]*entity.UserAttachedCoupon, error) {
	ctx := context.Background()
	pager := types.NewPageQuery(nil, &limit)
	userCoupons, err := models.CouponAttachedUsers(
		models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		models.CouponAttachedUserWhere.UsedAt.IsNull(),
		qm.Load(models.CouponAttachedUserRels.Coupon), // リレーションをロード
		qm.InnerJoin(models.TableNames.Coupon+" ON "+models.TableNames.CouponAttachedUser+"."+models.CouponAttachedUserColumns.CouponID+" = "+models.TableNames.Coupon+".id"),
		models.CouponWhere.ExpireAt.LT(time.Now()), // 親テーブルの条件を指定
		qm.Limit(pager.Limit()),
		qm.Offset(pager.Offset()),
		qm.OrderBy(models.CouponAttachedUserColumns.CouponID+" DESC"),
	).All(ctx, pq.db)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if userCoupons == nil {
		return nil, nil
	}
	var result = ModelToUserAttachedCoupon(userID, &userCoupons)
	return result, nil
}

func (pq *UserCouponQueryService) GetUseds(userID uuid.UUID, limit int) ([]*entity.UserAttachedCoupon, error) {
	pager := types.NewPageQuery(nil, &limit)
	userCoupons, err := models.CouponAttachedUsers(
		models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		models.CouponAttachedUserWhere.UsedAt.IsNotNull(),
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset()),
		qm.OrderBy(models.CouponAttachedUserColumns.CouponID+" DESC"),
	).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if userCoupons == nil {
		return nil, nil
	}
	var result = ModelToUserAttachedCoupon(userID, &userCoupons)
	return result, nil
}

func (pq *UserCouponQueryService) GetAll(userID uuid.UUID, pager *types.PageQuery) ([]*entity.UserAttachedCoupon, *types.PageResponse, error) {
	userCoupons, err := models.CouponAttachedUsers(models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon),
		qm.Limit(pager.Limit()), qm.Offset(pager.Offset()),
		qm.OrderBy(models.CouponAttachedUserColumns.CouponID+" DESC"),
	).All(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		logger.Error(err.Error())
		return nil, nil, nil
	}
	if userCoupons == nil {
		return nil, nil, nil
	}

	count, err := models.CouponAttachedUsers(models.CouponAttachedUserWhere.UserID.EQ(userID.String()),
		qm.Load(models.CouponAttachedUserRels.Coupon)).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	var pageResponse *types.PageResponse = nil
	if pager != nil {
		pageResponse = types.NewPageResponse(pager, int(count))
	}

	var result = ModelToUserAttachedCoupon(userID, &userCoupons)
	return result, pageResponse, err
}

func ModelToUserAttachedCoupon(userID uuid.UUID, userCoupons *models.CouponAttachedUserSlice) []*entity.UserAttachedCoupon {
	var result []*entity.UserAttachedCoupon

	for _, userCoupon := range *userCoupons {
		coupon := userCoupon.R.Coupon
		entityCoupon := CouponModelToEntity(coupon, nil)
		entityUserCoupon := entity.RegenUserAttachedCoupon(
			userID,
			entityCoupon,
			userCoupon.UsedAt.Ptr(),
		)
		result = append(result, entityUserCoupon)
	}
	return result
}
