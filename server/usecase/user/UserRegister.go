package user

import (
	"server/domain"
	"server/infra/interfaces"
	"time"

	"errors"

	"github.com/google/uuid"
)

type UserRegisterUsecase struct {
	userRepository interfaces.IUserRepository
}

func (u *UserRegisterUsecase) New(userRepository interfaces.IUserRepository) *UserRegisterUsecase {
	return &UserRegisterUsecase{
		userRepository: userRepository,
	}
}

func (u *UserRegisterUsecase) Exec(
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
	AcceptMail bool,
) (*domain.User, error) {

	existUser, err := u.userRepository.FindByMail(Mail)
	if err != nil {
		return nil, errors.New("ユーザーの検索に失敗しました")
	}

	if existUser != nil {
		return nil, errors.New("既に登録されているメールアドレスです")
	}

	user := (&domain.User{}).New(
		uuid.New(),
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

	return user, err
}
