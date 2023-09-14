package admin

import (
	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/core/infra/types"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	adminRepository repository.IAdminRepository
	adminQuery      queryservice.IAdminQueryService
	storeQuery      queryservice.IStoreQueryService
	authAction      action.IAuthAction
}

func NewAuthUsecase(adminRepository repository.IAdminRepository, adminQuery queryservice.IAdminQueryService,
	storeQuery queryservice.IStoreQueryService, authAction action.IAuthAction) *AuthUsecase {
	return &AuthUsecase{
		adminRepository: adminRepository,
		adminQuery:      adminQuery,
		storeQuery:      storeQuery,
		authAction:      authAction,
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
	newID, err := u.authAction.InviteUserByEmail(email)

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
		newID,
		name,
		email,
		true,
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
) error {
	err := u.authAction.SignUp(Mail, Password)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *AuthUsecase) SignIn(
	Mail string,
	Password string,
) (*types.Token, *errors.DomainError) {
	existUser, err := u.adminQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}
	token, err := u.authAction.SignIn(Mail, Password)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return token, nil
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
