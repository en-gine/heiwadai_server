package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/core/infra/types"
	"server/infrastructure/env"
	"server/infrastructure/logger"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	adminRepository        repository.IAdminRepository
	adminQuery             queryservice.IAdminQueryService
	userQuery              queryservice.IUserQueryService
	storeQuery             queryservice.IStoreQueryService
	userLoginLogRepository repository.IUserLoginLogRepository
	authAction             action.IAuthAction
}

func NewAuthUsecase(
	adminRepository repository.IAdminRepository,
	adminQuery queryservice.IAdminQueryService,
	userQuery queryservice.IUserQueryService,
	storeQuery queryservice.IStoreQueryService,
	userLoginLogRepository repository.IUserLoginLogRepository,
	authAction action.IAuthAction,
) *AuthUsecase {
	return &AuthUsecase{
		adminRepository:        adminRepository,
		adminQuery:             adminQuery,
		userQuery:              userQuery,
		storeQuery:             storeQuery,
		userLoginLogRepository: userLoginLogRepository,
		authAction:             authAction,
	}
}

func (u *AuthUsecase) Register(
	name string,
	storeID uuid.UUID,
	email string,
) (*entity.Admin, *errors.DomainError) {
	existAdmin, err := u.adminQuery.GetByMail(email)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "管理者ユーザーの検索に失敗しました")
	}

	if existAdmin != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既に登録されているメールアドレスです")
	}

	existUser, err := u.userQuery.GetByMail(email)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既にアプリユーザーとして登録されているメールアドレスです")
	}

	// 招待メール送信
	defaultPass := env.GetEnv(env.TestUserPass)
	pass, domainErr := entity.NewPassword(defaultPass)
	if domainErr != nil {
		return nil, domainErr
	}

	newID, err := u.authAction.SignUp(email, *pass)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	belongStore, err := u.storeQuery.GetByID(storeID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "店舗の検索に失敗しました")
	}

	if belongStore == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "IDで指定された店舗が見つかりません")
	}

	adminData := entity.RegenAdmin(
		*newID,
		name,
		email,
		true,
		false,
		belongStore,
	)

	err = u.adminRepository.Insert(adminData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return adminData, nil
}

func (u *AuthUsecase) SignUp(
	Mail string,
	Password string,
) (*uuid.UUID, *errors.DomainError) {

	pass, domainErr := entity.NewPassword(Password)
	if domainErr != nil {
		return nil, domainErr
	}

	id, err := u.authAction.SignUp(Mail, *pass)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return id, nil
}

func (u *AuthUsecase) SignOut(
	token string,
) *errors.DomainError {
	err := u.authAction.SignOut(token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) SignIn(
	Mail string,
	Password string,
	RemoteIP string,
	UserAgent string,
) (*types.Token, *errors.DomainError) {
	existUser, err := u.adminQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}

	if !existUser.IsActive {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "このアドレスで登録されているユーザーは無効化されています")
	}

	pass, domainErr := entity.NewPassword(Password)
	if domainErr != nil {
		return nil, domainErr
	}

	token, err := u.authAction.SignIn(Mail, *pass)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	loginLog := entity.CreateUserLoginLog(existUser.ID, RemoteIP, UserAgent)
	go func() {
		// ログイン履歴を保存
		err := u.userLoginLogRepository.Save(loginLog)
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	return token, nil
}

func (u *AuthUsecase) ReInviteMail(
	Mail string,
) *errors.DomainError {
	_, err := u.authAction.InviteUserByEmail(Mail)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) ResetPasswordMail(
	Mail string,
) *errors.DomainError {
	err := u.authAction.ResetPasswordMail(Mail)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdatePassword(
	Password string,
	Token string,
) *errors.DomainError {
	pass, domainErr := entity.NewPassword(Password)
	if domainErr != nil {
		return domainErr
	}
	err := u.authAction.UpdatePassword(*pass, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdateEmail(
	Mail string,
	Token string,
) *errors.DomainError {
	err := u.authAction.UpdateEmail(Mail, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) Refresh(
	Token string,
	RefreshToken string,
) (*types.Token, *errors.DomainError) {
	auth, err := u.authAction.Refresh(Token, RefreshToken)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	if auth == nil {
		return nil, errors.NewDomainError(errors.RepositoryError, "トークンの取得に失敗しました")
	}
	return auth.Token, nil
}
