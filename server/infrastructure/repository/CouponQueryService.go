package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

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
	coupon, err := models.FindCoupon(context.Background(), pq.db, id.String())
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	sotres, err := coupon.CouponStore(qm.Load(models.CouponStoreRels.Store)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var TargetStores []*entity.Store
	for _, store := range sotres {
		TargetStores = append(TargetStores, StoreModelToEntity(store.R.Store))
	}

	notices, err := coupon.CouponNotice().All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}

	var noticeResult []string
	for _, notice := range notices {
		noticeResult = append(noticeResult, notice.Notice)
	}

	return CouponModelToEntity(coupon, noticeResult, TargetStores), nil
}

func (pq *CouponQueryService) GetCouponListByType(couponType entity.CouponType, pager *types.PageQuery) ([]*entity.Coupon, error) {
	coupons, err := models.Coupons(models.CouponWhere.CouponType.EQ(couponType.ToInt()), qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	var result []*entity.Coupon
	for _, coupon := range coupons {
		result = append(result, CouponModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func (pq *CouponQueryService) GetCouponByType(couponType entity.CouponType) (*entity.Coupon, error) {
	coupon, err := models.Coupons(models.CouponWhere.CouponType.EQ(couponType.ToInt())).One(context.Background(), pq.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	sotres, err := coupon.CouponStore(qm.Load(models.CouponStoreRels.Store)).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var TargetStores []*entity.Store
	for _, store := range sotres {
		TargetStores = append(TargetStores, StoreModelToEntity(store.R.Store))
	}

	notices, err := coupon.CouponNotice().All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var noticeResult []string
	for _, notice := range notices {
		noticeResult = append(noticeResult, notice.Notice)
	}

	return CouponModelToEntity(coupon, noticeResult, TargetStores), nil
}

func CouponModelToEntity(model *models.Coupon, notices []string, targetStore []*entity.Store) *entity.Coupon {
	return entity.RegenCoupon(
		uuid.MustParse(model.ID),
		model.Name,
		entity.CouponType(model.CouponType),
		uint(model.DiscountAmount),
		model.ExpireAt,
		model.IsCombinationable,
		notices,
		targetStore,
		model.CreateAt,
		entity.CouponStatus(model.CouponStatus),
	)
}
