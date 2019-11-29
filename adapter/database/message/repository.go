package message

import (
	"github.com/jinzhu/gorm"

	"gin-gorm-sample/adapter/database/user"
	"gin-gorm-sample/application/domain"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) GetAll(userID uint) ([]domain.Message, error) {
	var messages []Message
	var domainMessages []domain.Message

	sql := repository.db.Preload("User")

	if userID != 0 {
		sql = sql.Where("user_id = ?", userID)
	}

	if err := sql.Find(&messages).Error; err != nil {
		return nil, err
	}

	for _, message := range messages {
		domainMessages = append(domainMessages, ConvertToDomainMessage(message))
	}

	return domainMessages, nil
}

func (repository *Repository) Get(messageID uint) (domain.Message, error) {
	var message Message

	err := repository.db.Preload("User").First(&message, messageID).Error

	return ConvertToDomainMessage(message), err
}

func (repository *Repository) Create(domainMessage *domain.Message) error {
	message := Message{
		Context: domainMessage.Context,
		UserID:  domainMessage.UserID,
	}

	return repository.db.Create(&message).Error
}

func (repository *Repository) Update(domainMessage *domain.Message) error {
	message := Message{
		ID:      domainMessage.ID,
		Context: domainMessage.Context,
		UserID:  domainMessage.UserID,
	}

	return repository.db.Model(&message).Updates(message).Error
}

func (repository *Repository) Delete(id uint) error {
	message := Message{
		ID: id,
	}
	return repository.db.Delete(&message).Error
}

func ConvertToDomainMessage(message Message) domain.Message {
	domainMessage := domain.Message{
		ID:        message.ID,
		Context:   message.Context,
		User:      user.ConvertToDomainUser(message.User),
		UpdatedAt: message.UpdatedAt,
		CreatedAt: message.CreatedAt,
	}
	return domainMessage
}
