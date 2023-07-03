package user

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"time"

	"github.com/google/uuid"
)

type UserDataUsecase struct {
	userRepository repository.IUserRepository
	userQuery      queryservice.IUserQueryService
}

func NewUserDataUsecase(userRepository repository.IUserRepository, userQuery queryservice.IUserQueryService) *UserDataUsecase {
	return &UserDataUsecase{
		userRepository: userRepository,
	}
}

func (u *UserDataUsecase) Create(
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture string,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool, // メルマガ配信可
) (*entity.User, *errors.DomainError) {

	existUser, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.NewDomainError(errors.UnPemitedOperation, "既に登録されているメールアドレスです")
	}

	insertData := entity.CreateUser(
		FirstName,
		LastName,
		FirstNameKana,
		LastNameKana,
		CompanyName,
		BirthDate,
		ZipCode,
		Prefecture,
		City,
		Address,
		Tel,
		Mail,
		AcceptMail,
	)
	err = u.userRepository.Save(insertData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}
	return insertData, nil
}

func (u *UserDataUsecase) Update(
	ID uuid.UUID,
	FirstName string,
	LastName string,
	FirstNameKana string,
	LastNameKana string,
	CompanyName *string,
	BirthDate time.Time,
	ZipCode *string,
	Prefecture string,
	City *string,
	Address *string,
	Tel *string,
	Mail string,
	AcceptMail bool, // メルマガ配信可
) (*entity.User, *errors.DomainError) {

	existUser, err := u.userQuery.GetById(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, "ユーザーの検索に失敗しました")
	}

	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
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
		Prefecture,
		City,
		Address,
		Tel,
		Mail,
		AcceptMail,
	)
	err = u.userRepository.Save(updateData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateData, nil
}
