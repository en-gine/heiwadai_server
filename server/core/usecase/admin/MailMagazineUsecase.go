package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"
	"server/infrastructure/logger"

	"github.com/google/uuid"
)

type MailMagazineUsecase struct {
	mailMagazineRepository    repository.IMailMagazineRepository
	mailMagazineQuery         queryservice.IMailMagazineQueryService
	mailMagazineLogQuery      queryservice.IMailMagazineLogQueryService
	mailMagazineLogRepository repository.IMailMagazineLogRepository
	mailSendAction            action.ISendMailAction
	userQueryService          queryservice.IUserQueryService
}

func NewMailMagazineUsecase(
	mailMagazineRepository repository.IMailMagazineRepository,
	mailMagazineQuery queryservice.IMailMagazineQueryService,
	mailMagazineLogQuery queryservice.IMailMagazineLogQueryService,
	mailMagazineLogRepository repository.IMailMagazineLogRepository,
	mailSendAction action.ISendMailAction,
	userQueryService queryservice.IUserQueryService,
) *MailMagazineUsecase {
	return &MailMagazineUsecase{
		mailMagazineRepository:    mailMagazineRepository,
		mailMagazineQuery:         mailMagazineQuery,
		mailMagazineLogRepository: mailMagazineLogRepository,
		mailMagazineLogQuery:      mailMagazineLogQuery,
		mailSendAction:            mailSendAction,
		userQueryService:          userQueryService,
	}
}

func (u *MailMagazineUsecase) GetList(pager *types.PageQuery) ([]*entity.MailMagazine, *types.PageResponse, *errors.DomainError) {
	mailMagazines, page, err := u.mailMagazineQuery.GetAll(pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return mailMagazines, page, nil
}

func (u *MailMagazineUsecase) GetLogList(userID uuid.UUID, pager types.PageQuery) ([]*entity.MailMagazineLogWithTitle, *types.PageResponse, *errors.DomainError) {
	mailMagazines, pageRes, err := u.mailMagazineLogQuery.GetUserLogList(userID, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return mailMagazines, pageRes, nil
}

func (u *MailMagazineUsecase) SaveDraft(title string, content string, targetPrefectures *[]entity.Prefecture, autherID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
	maySendCount, err := u.userQueryService.GetMailOKUserCount(targetPrefectures)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	mailMagazine := entity.CreateSavedMailMagazine(title, content, targetPrefectures, *maySendCount, autherID)

	err = u.mailMagazineRepository.Save(mailMagazine)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return mailMagazine, nil
}

func (u *MailMagazineUsecase) Update(title *string, content *string, targetPrefectures *[]entity.Prefecture, autherID uuid.UUID, mailMagazineID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
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

	maySendCount, err := u.userQueryService.GetMailOKUserCount(targetPrefectures)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	updateMailMagazine := entity.UpdateMailMagazine(mailMagazineID, updateTitle, updateContent, targetPrefectures, *maySendCount, autherID)

	if updateMailMagazine.MailMagazineStatus == entity.MailMagazineSentUnCompleted {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "送信処理済みのため編集できません")
	}

	if updateMailMagazine.MailMagazineStatus == entity.MailMagazineSentCompleted {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "送信済みのため編集できません")
	}

	err = u.mailMagazineRepository.Save(updateMailMagazine)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateMailMagazine, nil
}

func (u *MailMagazineUsecase) GetByID(mailMagazineID uuid.UUID) (*entity.MailMagazine, *errors.DomainError) {
	mgz, err := u.mailMagazineQuery.GetByID(mailMagazineID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return mgz, nil
}

func (u *MailMagazineUsecase) Delete(mailMagazineID uuid.UUID) *errors.DomainError {
	deleteMailMagazine, err := u.mailMagazineQuery.GetByID(mailMagazineID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if deleteMailMagazine == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}

	if deleteMailMagazine.MailMagazineStatus == entity.MailMagazineSentCompleted || deleteMailMagazine.MailMagazineStatus == entity.MailMagazineSentUnCompleted {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "送信処理済みのため削除できません")
	}

	err = u.mailMagazineRepository.Delete(mailMagazineID)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}

func (u *MailMagazineUsecase) Send(mailMagazineID uuid.UUID) *errors.DomainError {
	mgz, err := u.mailMagazineQuery.GetByID(mailMagazineID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	if mgz == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "対象の投稿が見つかりません")
	}

	if mgz.MailMagazineStatus == entity.MailMagazineDraft {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "下書きは一度保存した後に送信できます")
	}

	if mgz.MailMagazineStatus == entity.MailMagazineSentCompleted {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "送信済みのため送信できません")
	}

	// 一旦ログにコピーしてDB保存（途中失敗なども考えて状態を固定）／送信途中失敗時はコピーしない
	if mgz.MailMagazineStatus != entity.MailMagazineSentUnCompleted {
		err = u.mailMagazineLogRepository.BulkCopyToLogAsUnsent(mailMagazineID, mgz.TargetPrefecture)
	}

	if err != nil {
		u.saveUncompleteMailMagazine(mgz, mgz.UnsentCount, mgz.SentCount)
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	count, err := u.mailMagazineLogQuery.GetUnsentTargetAllCount(mailMagazineID)
	if err != nil {
		u.saveUncompleteMailMagazine(mgz, mgz.UnsentCount, mgz.SentCount)
		return errors.NewDomainError(errors.QueryError, err.Error())
	}

	PerPage := 1000
	pages := count / PerPage
	if count%PerPage != 0 {
		pages++ // 余りがある場合はページを追加
	}
	atFirstUnsent := mgz.UnsentCount
	for page := 1; page <= pages; page++ {
		pager := types.NewPageQuery(
			&page,
			&PerPage,
		)

		prevSend := (page - 1) * 1000 // 一つ前のループまでの送信完了数

		unsentMails, err := u.mailMagazineLogQuery.GetUnsentTargetMails(mailMagazineID, *pager)
		if err != nil {
			u.saveUncompleteMailMagazine(mgz, atFirstUnsent-prevSend, prevSend)
			return errors.NewDomainError(errors.QueryError, err.Error())
		}
		// 送信処理
		err = u.mailSendAction.SendAll(unsentMails, "no-reply@heiwadai-hotel.app", mgz.Title, mgz.Content)
		if err != nil {
			u.saveUncompleteMailMagazine(mgz, atFirstUnsent-prevSend, prevSend)
			return errors.NewDomainError(errors.RepositoryError, err.Error())
		}
		// 送信済みに更新
		err = u.mailMagazineLogRepository.BulkMarkAsSent(mailMagazineID, *pager)
		if err != nil {
			sendCount := page * 1000 // 送信完了予定数
			if sendCount > count {
				sendCount = count
			}
			u.saveUncompleteMailMagazine(mgz, atFirstUnsent-prevSend, sendCount)
			return errors.NewDomainError(errors.RepositoryError, err.Error())
		}
	}

	completeCount := mgz.UnsentCount + mgz.SentCount
	completeMagazine := entity.CreateSentCompleteMailMagazine(mgz.ID, mgz.Title, mgz.TargetPrefecture, completeCount, mgz.Content, mgz.AuthorID)
	err = u.mailMagazineRepository.Save(completeMagazine)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return nil
}

// 処理が失敗した時にメールマガジンのステータスをUnsentにして保存する処理
func (u *MailMagazineUsecase) saveUncompleteMailMagazine(mgz *entity.MailMagazine, unsentCount int, sentCount int) {
	uncompleteMgz := entity.CreateUnCompleteMailMagazine(mgz.ID, mgz.Title, mgz.TargetPrefecture, mgz.UnsentCount, mgz.SentCount, mgz.Content, mgz.AuthorID)
	err := u.mailMagazineRepository.Save(uncompleteMgz)
	if err != nil {
		logger.Error(err.Error())
	}
}
