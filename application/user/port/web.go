package port

import "gin-gorm-sample/application/domain"

type Web interface {
	Get(userID uint) (domain.User, error)
	GetAll() ([]domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error
}
