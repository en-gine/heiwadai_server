package repository

import (
	"context"
	"database/sql"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"
	"server/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.ICouponQueryService = &CouponQueryService{}

type CouponQueryService struct {
	db *sql.DB
}

func NewCouponQueryService() *CouponQueryService {
	db := InitDB()

	return &CouponQueryService{
		db: db,
	}
}

func (pq *CouponQueryService) GetByID(id uuid.UUID) (*entity.Coupon, error) {
	coupon, err := models.Coupons(
		models.CouponWhere.ID.EQ(id.String()),
		qm.Load(models.CouponRels.Stores),
	).One(context.Background(), pq.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}
	if coupon == nil {
		return nil, nil
	}
	stores := coupon.R.Stores
	var TargetStores []*entity.Store
	for _, store := range stores {
		TargetStores = append(TargetStores, StoreModelToEntity(store, nil))
	}

	return CouponModelToEntity(coupon, TargetStores), nil
}

func (pq *CouponQueryService) GetCouponListByType(couponType entity.CouponType, pager *types.PageQuery) ([]*entity.Coupon, *types.PageResponse, error) {
	coupons, err := models.Coupons(models.CouponWhere.CouponType.EQ(couponType.ToInt()), qm.Limit(pager.Limit()), qm.Offset(pager.Offset()), qm.OrderBy(models.CouponColumns.CreateAt+" DESC")).All(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}
	var result []*entity.Coupon
	for _, coupon := range coupons {
		result = append(result, CouponModelToEntity(coupon, nil))
	}

	count, err := models.Coupons(models.CouponWhere.CouponType.EQ(couponType.ToInt())).Count(context.Background(), pq.db)
	if err != nil {
		return nil, nil, err
	}

	pageResponse := types.NewPageResponse(pager, int(count))

	return result, pageResponse, err
}

func (pq *CouponQueryService) GetCouponByType(couponType entity.CouponType) (*entity.Coupon, error) {
	coupon, err := models.Coupons(models.CouponWhere.CouponType.EQ(couponType.ToInt()), qm.Load(models.CouponRels.Stores)).One(context.Background(), pq.db)
	if coupon == nil {
		return nil, nil
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, nil
	}

	stores := coupon.R.Stores

	var TargetStores []*entity.Store
	for _, store := range stores {
		TargetStores = append(TargetStores, StoreModelToEntity(store, nil))
	}

	return CouponModelToEntity(coupon, TargetStores), nil
}

func CouponModelToEntity(model *models.Coupon, targetStore []*entity.Store) *entity.Coupon {
	return entity.RegenCoupon(
		uuid.MustParse(model.ID),
		model.Name,
		entity.CouponType(model.CouponType),
		uint(model.DiscountAmount),
		model.ExpireAt,
		model.IsCombinationable,
		model.Notices,
		targetStore,
		model.CreateAt,
		entity.CouponStatus(model.CouponStatus),
		&model.IssueCount,
		model.IssueAt.Ptr(),
	)
}
