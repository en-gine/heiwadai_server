package user

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	"server/core/infra/action"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/core/infra/types"
	"server/infrastructure/logger"

	"github.com/google/uuid"
)

type AuthUsecase struct {
	userRepository         repository.IUserRepository
	userQuery              queryservice.IUserQueryService
	userLoginLogRepository repository.IUserLoginLogRepository
	authAction             action.IAuthAction
}

func NewAuthUsecase(
	userRepository repository.IUserRepository,
	userQuery queryservice.IUserQueryService,
	userLoginLogRepository repository.IUserLoginLogRepository,
	authAction action.IAuthAction,
) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		userQuery:      userQuery,
		authAction:     authAction,
	}
}

func (u *AuthUsecase) Register(
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate *time.Time,
	ZipCode *string,
	PrefectureID int,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool, // メルマガ配信可
	AcceptTerm bool, // 利用規約に同意
) (*entity.User, *errors.DomainError) {
	if !AcceptTerm {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "利用規約に同意してください")
	}

	existUser, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既に登録されているメールアドレスです")
	}

	prefecture := entity.Prefecture(PrefectureID)

	// 招待メール送信
	newID, err := u.authAction.InviteUserByEmail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	userData := entity.CreateUser(
		*newID,
		FirstName,
		LastName,
		FirstNameKana,
		LastNameKana,
		CompanyName,
		BirthDate,
		ZipCode,
		prefecture,
		City,
		Address,
		Tel,
		Mail,
		AcceptMail,
	)

	err = u.userRepository.Save(userData, nil)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return userData, nil
}

func (u *AuthUsecase) SignUp(
	Mail string,
	Password string,
) (*uuid.UUID, error) {
	userID, err := u.authAction.SignUp(Mail, Password, action.UserTypeUser)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return userID, nil
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
	existUser, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}

	userOption, err := u.userQuery.GetOptionByID(existUser.ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if userOption.IsBlackCustomer {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "このアドレスで登録されているユーザーは無効化されています")
	}

	token, err := u.authAction.SignIn(Mail, Password)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	// ログイン履歴を保存敢えてエラーは無視
	loginLog := entity.CreateUserLoginLog(existUser.ID, RemoteIP, UserAgent)
	err = u.userLoginLogRepository.Save(loginLog)
	if err != nil {
		logger.Errorf("ログイン履歴の保存に失敗しました: %s", err.Error())
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
