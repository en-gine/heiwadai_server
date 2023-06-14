package mocks

import (
	"server/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// UserRepository は domain.UserRepository のモックを表す構造体です。
type UserRepository struct {
	mock.Mock
}

// FindByMail はメールアドレスからユーザーを検索するメソッドのモックです。
func (m *UserRepository) FindByMail(mail string) (*domain.User, error) {
	args := m.Called(mail)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Create は新しいユーザーを作成するメソッドのモックです。
func (m *UserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepository) FindById(uuid uuid.UUID) (*domain.User, error) {
	args := m.Called(uuid)
	return args.Get(0).(*domain.User), args.Error(1)
}
