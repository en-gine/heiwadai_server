package interfaces

import "server/domain"

type IUserRepository interface {
	Save(user *domain.User) error
	Update(user *domain.User) error
	FindById() (*domain.User, error)
	FindByMail(mail string) (*domain.User, error)
}
