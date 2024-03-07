package cron

import (
	"context"
	"fmt"
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
)

type CronCouponUsecase struct {
	couponRepository       repository.ICouponRepository
	couponQuery            queryservice.ICouponQueryService
	userCouponQuery        queryservice.IUserCouponQueryService
	usercouponRepository   repository.IUserCouponRepository
	storeQuery             queryservice.IStoreQueryService
	cronIssueLogRepository repository.ICronIssueLogRepository
	cronIssueLogQuery      queryservice.ICronIssueLogQueryService
	transaction            repository.ITransaction
}

func NewCronCouponUsecase(couponRepository repository.ICouponRepository, couponQuery queryservice.ICouponQueryService,
	userCouponQuery queryservice.IUserCouponQueryService, usercouponRepository repository.IUserCouponRepository, storeQuery queryservice.IStoreQueryService,
	cronIssueLogRepository repository.ICronIssueLogRepository, cronIssueLogQuery queryservice.ICronIssueLogQueryService,
	transaction repository.ITransaction,
) *CronCouponUsecase {
	return &CronCouponUsecase{
		couponRepository:       couponRepository,
		couponQuery:            couponQuery,
		userCouponQuery:        userCouponQuery,
		usercouponRepository:   usercouponRepository,
		storeQuery:             storeQuery,
		cronIssueLogRepository: cronIssueLogRepository,
		cronIssueLogQuery:      cronIssueLogQuery,
		transaction:            transaction,
	}
}

func (u *CronCouponUsecase) BulkAttachBirthdayCoupon(birthMonth int) (*int, *errors.DomainError) {
	ctx := context.Background()

	currentYear := time.Now().Year()
	hasIssued, err := u.cronIssueLogQuery.HasYearMonthLog(currentYear, birthMonth)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if hasIssued {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, fmt.Sprintf("既に%d年%dのお誕生日クーポンは発行済みです", currentYear, birthMonth))
	}

	err = u.transaction.Begin(ctx)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	allStores, err := u.storeQuery.GetActiveAll()
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	birthdayCoupon, domainErr := entity.CreateBirthdayCoupon(allStores)
	if domainErr != nil {
		u.transaction.Rollback()
		return nil, domainErr
	}

	err = u.couponRepository.Save(u.transaction, birthdayCoupon)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	count, err := u.usercouponRepository.IssueAll(u.transaction, birthdayCoupon, &birthMonth)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.ActionError, err.Error())
	}

	issuedCoupon := entity.CreateIssuedCoupon(birthdayCoupon, &count)
	err = u.couponRepository.Save(u.transaction, issuedCoupon)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	err = u.cronIssueLogRepository.Save(u.transaction, birthdayCoupon.Name+"自動発行", count, currentYear, birthMonth)
	if err != nil {
		u.transaction.Rollback()
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	err = u.transaction.Commit()
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return &count, nil
}
