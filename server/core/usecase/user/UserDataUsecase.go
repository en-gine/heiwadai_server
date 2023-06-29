package user

import (
	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"time"

	"errors"

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
) (*entity.User, error) {

	existUser, err := u.userQuery.GetByMail(Mail)
	if err != nil {
		return nil, errors.New("ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.New("既に登録されているメールアドレスです")
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
	return u.userRepository.Save(insertData)
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
) (*entity.User, error) {

	existUser, err := u.userQuery.GetById(ID)
	if err != nil {
		return nil, errors.New("ユーザーの検索に失敗しました")
	}

	if existUser == nil {
		return nil, errors.New("登録されているユーザーが存在しません")
	}

	updateData := entity.StoredUser(
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
	return u.userRepository.Save(updateData)
}
