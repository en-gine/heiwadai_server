package user

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserReportUsecase struct {
	userQuery            queryservice.IUserQueryService
	userReportRepository repository.IUserReportRepository
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

func (u *UserReportUsecase) Send(title string, content string, userID uuid.UUID) *errors.DomainError {
	user, err := u.userQuery.GetByID(userID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if user == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "対象のユーザーが見つかりません")
	}

	userName := user.LastName + user.FirstName
	report := entity.CreateUserReport(title, content, userID, userName)

	err = u.userReportRepository.Save(report)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	mailContent := userName + " 様より下記の内容でお問い合わせがありました。\n\n" + report.Content
	err = u.sendMailAction.Send(user.Mail, report.Title, mailContent, action.SendStylePlainText)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	return nil
}
