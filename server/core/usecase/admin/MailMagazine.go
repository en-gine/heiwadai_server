package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type MailMagazineUsecase struct {
	mailMagazineRepository repository.IMailMagazineRepository
	mailMagazineQuery      queryservice.IMailMagazineQueryService
}

func NewMailMagazineUsecase(mailMagazineRepository repository.IMailMagazineRepository, mailMagazineQuery queryservice.IMailMagazineQueryService) *MailMagazineUsecase {
	return &MailMagazineUsecase{
		mailMagazineRepository: mailMagazineRepository,
		mailMagazineQuery:      mailMagazineQuery,
	}
}

func (u *MailMagazineUsecase) GetList(pager *types.PageQuery) ([]*entity.MailMagazine, *errors.DomainError) {
	mailMagazines, err := u.mailMagazineQuery.GetAll(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return mailMagazines, nil
}

func (u *MailMagazineUsecase) Create(title string, content string, autherID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
	mailMagazine := entity.CreateDraftMailMagazine(title, content, autherID)

	err := u.mailMagazineRepository.Save(mailMagazine)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return mailMagazine, nil
}

func (u *MailMagazineUsecase) Update(title *string, content *string, autherID uuid.UUID, mailMagazineID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
	oldMailMagazine, err := u.mailMagazineQuery.GetByID(mailMagazineID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if oldMailMagazine == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}
	var updateTitle string
	var updateContent string
	if title != nil {
		updateTitle = *title
	} else {
		updateTitle = oldMailMagazine.Title
	}

	if content != nil {
		updateContent = *content
	} else {
		updateContent = oldMailMagazine.Content
	}

	updateMailMagazine := entity.UpdateMailMagazine(mailMagazineID, updateTitle, updateContent, autherID)

	if updateMailMagazine.MailMagazineStatus == entity.MailMagazineSent {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "送信済みのため編集できません")
	}

	err = u.mailMagazineRepository.Save(updateMailMagazine)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateMailMagazine, nil
}

func (u *MailMagazineUsecase) Delete(mailMagazineID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
	deleteMailMagazine, err := u.mailMagazineQuery.GetByID(mailMagazineID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if deleteMailMagazine == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}

	if deleteMailMagazine.MailMagazineStatus == entity.MailMagazineSent {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "送信済みのため削除できません")
	}

	err = u.mailMagazineRepository.Delete(mailMagazineID)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return deleteMailMagazine, nil
}
