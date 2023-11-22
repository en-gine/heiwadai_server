package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserCheckinUsecase struct {
	checkInRepository repository.ICheckinRepository
	userCheckinQuery  queryservice.ICheckinQueryService
}

func NewUserCheckinUsecase(
	checkInRepository repository.ICheckinRepository,
) *UserCheckinUsecase {
	return &UserCheckinUsecase{
		checkInRepository: checkInRepository,
	}
}

func (u *UserCheckinUsecase) GetAllRecent(limit int) ([]*entity.Checkin, *errors.DomainError) {
	pager := &types.PageQuery{
		PerPage: &limit,
	}
	checkins, err := u.userCheckinQuery.GetAllUserAllCheckin(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return checkins, nil
}

func (u *UserCheckinUsecase) GetUserLog(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, *types.PageResponse, *errors.DomainError) {
	checkins, pageResponse, err := u.userCheckinQuery.GetMyAllCheckin(userID, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return checkins, pageResponse, nil
}
