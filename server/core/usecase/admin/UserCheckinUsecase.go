package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"
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

func (u *UserCheckinUsecase) GetRecent() ([]*entity.Checkin, *errors.DomainError) {
	perPage := 30
	pager := &types.PageQuery{
		PerPage: &perPage,
	}
	checkins, err := u.userCheckinQuery.GetAllUserAllCheckin(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return checkins, nil
}
