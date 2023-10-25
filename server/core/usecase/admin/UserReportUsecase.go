package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserReportUsecase struct {
	userQuery            queryservice.IUserQueryService
	userReportRepository repository.IUserReportRepository
	userReportQuery      queryservice.IUserReportQueryService
	sendMailAction       action.ISendMailAction
}

func NewUserReportUsecase(
	userQuery queryservice.IUserQueryService,
	userReportRepository repository.IUserReportRepository,
	sendMailAction action.ISendMailAction,
) *UserReportUsecase {
	return &UserReportUsecase{
		userQuery:            userQuery,
		userReportRepository: userReportRepository,
		sendMailAction:       sendMailAction,
	}
}

func (u *UserReportUsecase) GetAll(pager *types.PageQuery) ([]*entity.UserReport, *errors.DomainError) {
	reports, err := u.userReportQuery.GetAll(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return reports, nil
}

func (u *UserReportUsecase) GetByID(id uuid.UUID) (*entity.UserReport, *errors.DomainError) {
	report, err := u.userReportQuery.GetByID(id)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return report, nil
}
