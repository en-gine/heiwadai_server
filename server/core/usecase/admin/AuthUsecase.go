package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/core/infra/types"
	"server/infrastructure/env"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	adminRepository        repository.IAdminRepository
	adminQuery             queryservice.IAdminQueryService
	storeQuery             queryservice.IStoreQueryService
	userLoginLogRepository repository.IUserLoginLogRepository
	authAction             action.IAuthAction
}

func NewAuthUsecase(
	adminRepository repository.IAdminRepository,
	adminQuery queryservice.IAdminQueryService,
	storeQuery queryservice.IStoreQueryService,
	userLoginLogRepository repository.IUserLoginLogRepository,
	authAction action.IAuthAction,
) *AuthUsecase {
	return &AuthUsecase{
		adminRepository:        adminRepository,
		adminQuery:             adminQuery,
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
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "ユーザーの検索に失敗しました")
	}

	if existAdmin != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既に登録されているメールアドレスです")
	}

	// 招待メール送信
	newID, err := u.authAction.SignUp(email, env.GetEnv(env.TestUserPass), action.UserTypeAdmin)
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
) (*uuid.UUID, error) {
	id, err := u.authAction.SignUp(Mail, Password, action.UserTypeUser)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return id, nil
}

func (u *AuthUsecase) SignOut(
	token string,
) error {
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

	token, err := u.authAction.SignIn(Mail, Password)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	loginLog := entity.CreateUserLoginLog(existUser.ID, RemoteIP, UserAgent)
	// ログイン履歴を保存敢えてエラーは無視
	_ = u.userLoginLogRepository.Save(loginLog)

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
) error {
	err := u.authAction.ResetPasswordMail(Mail)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdatePassword(
	Password string,
	Token string,
) error {
	err := u.authAction.UpdatePassword(Password, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) UpdateEmail(
	Mail string,
	Token string,
) error {
	err := u.authAction.UpdateEmail(Mail, Token)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) Refresh(
	Token string,
	RefreshToken string,
) (*types.Token, error) {
	auth, err := u.authAction.Refresh(Token, RefreshToken)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	if auth == nil {
		return nil, errors.NewDomainError(errors.RepositoryError, "トークンの取得に失敗しました")
	}
	return auth.Token, nil
}
