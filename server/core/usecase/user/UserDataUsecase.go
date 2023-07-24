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
		userQuery:      userQuery,
	}
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

	existUser, err := u.userQuery.GetByID(ID)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	if existUser == nil {
		return nil, errors.NewDomainError(errors.QueryDataNotFoundError, "登録されているユーザーが存在しません")
	}
	prefecture, domainErr := entity.StringToPrefecture(Prefecture)
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
	err = u.userRepository.Save(updateData)
	if err != nil {
		return nil, errors.NewDomainError(errors.RepositoryError, err.Error())
	}

	return updateData, nil
}
