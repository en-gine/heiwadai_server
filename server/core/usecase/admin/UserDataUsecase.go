package admin

import (
	"time"

	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserDataUsecase struct {
	userRepository repository.IUserRepository
	userQuery      queryservice.IUserQueryService
}

func NewUserDataUsecase(userRepository repository.IUserRepository, userQuery queryservice.IUserQueryService) *UserDataUsecase {
	return &UserDataUsecase{
		userRepository: userRepository,
		userQuery:      userQuery,
	}
}

func (u *UserDataUsecase) GetList(query *types.UserQuery, pager *types.PageQuery) ([]*entity.UserWichLastCheckin, *types.PageResponse, *errors.DomainError) {
	users, page, err := u.userQuery.GetList(query, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return users, page, nil
}

func (u *UserDataUsecase) Update(
	ID uuid.UUID,
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
	InnerNoto string,
	IsBlackCustomer bool,
) (*entity.UserWithOption, *errors.DomainError) {
	existUser, err := u.userQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}

	existUser, err = u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "このアドレスで登録されているユーザーが存在しません")
	}

	prefecture, domainErr := entity.IntToPrefecture(PrefectureID)
	if domainErr != nil {
		return nil, domainErr
	}

	updateData := entity.RegenUser(
		ID,
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
	userOption := entity.CreateUserOption(ID, InnerNoto, IsBlackCustomer)

	err = u.userRepository.Save(updateData, userOption)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	userWithOption := &entity.UserWithOption{
		User:       updateData,
		UserOption: userOption,
	}
	return userWithOption, nil
}

func (u *UserDataUsecase) Delete(ID uuid.UUID) *errors.DomainError {
	user, err := u.userQuery.GetByID(ID)
	if err != nil {
		return errors.NewDomainError(errors.QueryError, err.Error())
	}
	if user == nil {
		return errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}

	err = u.userRepository.Delete(ID)
	if err != nil {
		return errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return nil
}

func (u *UserDataUsecase) GetUserByID(ID uuid.UUID) (*entity.UserWithOption, *errors.DomainError) {
	user, err := u.userQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	if user == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}
	option, err := u.userQuery.GetOptionByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return &entity.UserWithOption{
		User:       user,
		UserOption: option,
	}, nil
}
