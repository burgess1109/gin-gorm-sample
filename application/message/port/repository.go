package port

import (
	"gin-gorm-sample/application/domain"
)

type Repository interface {
	Get(messageID uint) (domain.Message, error)
	GetAll(userID uint) ([]domain.Message, error)
	Create(message *domain.Message) error
	Update(message *domain.Message) error
	Delete(id uint) error
}
