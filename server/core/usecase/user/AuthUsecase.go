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
	adminQuery             queryservice.IAdminQueryService
	userLoginLogRepository repository.IUserLoginLogRepository
	authAction             action.IAuthAction
}

func NewAuthUsecase(
	userRepository repository.IUserRepository,
	userQuery queryservice.IUserQueryService,
	adminQuery queryservice.IAdminQueryService,
	userLoginLogRepository repository.IUserLoginLogRepository,
	authAction action.IAuthAction,
) *AuthUsecase {
	return &AuthUsecase{
		userRepository:         userRepository,
		userQuery:              userQuery,
		adminQuery:             adminQuery,
		userLoginLogRepository: userLoginLogRepository,
		authAction:             authAction,
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
		return nil, errors.NewDomainError(errors.AlreadyExist, "既に登録されているメールアドレスです")
	}

	existAdmin, err := u.adminQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "管理者の検索に失敗しました")
	}

	if existAdmin != nil {
		return nil, errors.NewDomainError(errors.AlreadyExist, "既に内部スタッフとして登録されているメールアドレスです")
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
) (*uuid.UUID, *errors.DomainError) {
	pass, domainErr := entity.NewPassword(Password)

	if domainErr != nil {
		return nil, domainErr
	}

	userID, err := u.authAction.SignUp(Mail, *pass)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return userID, nil
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

	pass, domainErr := entity.NewPassword(Password)

	if domainErr != nil {
		return nil, domainErr
	}

	token, err := u.authAction.SignIn(Mail, *pass)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	// ログイン履歴を保存敢えてエラーは無視
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

func (u *AuthUsecase) GetUserByToken(
	Token string,
) (*string, *errors.DomainError) {
	info, err := u.authAction.GetUserInfo(Token)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return &info.Mail, nil
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

func (u *AuthUsecase) IsUnderRegister(
	Mail string,
) (bool, *errors.DomainError) {
	isExist, err := u.userQuery.IsUnderRegister(Mail)
	if err != nil {
		return false, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return isExist, nil
}

func (u *AuthUsecase) ResendInviteMail(
	Mail string,
) *errors.DomainError {
	_, err := u.authAction.InviteUserByEmail(Mail)

	return errors.NewDomainError(errors.ActionError, err.Error())
}

func (u *AuthUsecase) DeleteUnderRegisterUser(
	Mail string,
) *errors.DomainError {
	isExist, err := u.userQuery.IsUnderRegister(Mail)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if !isExist {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録中のユーザーが存在しません")
	}

	user, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if user == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスのユーザーが存在しません")
	}

	err = u.userRepository.Delete(user.ID)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}
