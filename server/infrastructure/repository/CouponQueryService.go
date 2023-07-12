package repository

import (
	"context"
	"database/sql"
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/db/models"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var _ queryservice.ICouponQueryService = &CouponQueryService{}

type CouponQueryService struct {
	db *sql.DB
}

func NewCouponQueryService() (*CouponQueryService, error) {
	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	return &CouponQueryService{
		db: db,
	}, nil
}

func (pq *CouponQueryService) GetById(id uuid.UUID) (*entity.Coupon, error) {
	coupon, err := models.FindCoupon(context.Background(), pq.db, id.String())
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

func (pq *CouponQueryService) GetActiveAll(user *entity.User) ([]*entity.Coupon, error) {
	coupons, err := models.Coupons(qm.Load(models.CouponRels.User), models.CouponWhere.CouponStatus.EQ(int(entity.CouponIssued)), models.CouponWhere.UserID.EQ(null.StringFrom(user.ID.String()))).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Coupon

	for _, coupon := range coupons {
		result = append(result, CouponModelToEntity(coupon, nil, nil))
	}
	return result, nil
}

func (pq *CouponQueryService) GetAll(user *entity.User, pager *types.PageQuery) ([]*entity.Coupon, error) {
	coupons, err := models.Coupons(models.CouponWhere.UserID.EQ(null.StringFrom(user.ID.String())), qm.Load(models.CouponRels.User), qm.Limit(pager.Offset()), qm.Offset(pager.Offset())).All(context.Background(), pq.db)
	if err != nil {
		return nil, err
	}
	var result []*entity.Coupon
	for _, coupon := range coupons {
		result = append(result, CouponModelToEntity(coupon, nil, nil))
	}
	return result, nil

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
		model.UsedAt.Ptr(),
		UserModelToEntity(model.R.User),
		targetStore,
		model.CreateAt,
		entity.CouponStatus(model.CouponStatus),
	)
}
