package user

import (
	"server/domain"
	"server/infra/interfaces"
	"time"

	"errors"

	"github.com/google/uuid"
)

type AuthCreateUsecase struct {
	userRepository interfaces.IUserRepository
}

func New(userRepository interfaces.IUserRepository) *AuthCreateUsecase {
	return &AuthCreateUsecase{
		userRepository: userRepository,
	}
}

func (u *AuthCreateUsecase) Exec(
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
