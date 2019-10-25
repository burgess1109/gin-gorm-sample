package message

import (
	"github.com/jinzhu/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	GetMessagesByUserID(userID uint) ([]MessageAssociateUser, error)
	GetMessageByID(messageID uint) (Message, error)
	CreateMessage(message *Message) error
	UpdateMessage(message *Message) error
	DeleteMessage(message *Message) error
}

type MessageAssociateUser struct {
	Message
	UserName string `json:"user_name"`
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return MessageRepository{db: db}
}

func (repository *MessageRepository) GetMessagesByUserID(userID uint) ([]MessageAssociateUser, error) {
	var messageAssociateUsers []MessageAssociateUser

	sqlDB := repository.db.
		Select("message.*, user.name as user_name").
		Joins("JOIN user ON message.user_id = user.id")

	if userID != 0 {
		sqlDB = sqlDB.Where("message.user_id = ?", userID)
	}

	err := sqlDB.Find(&messageAssociateUsers).Error

	return messageAssociateUsers, err
}

func (repository *MessageRepository) GetMessageByID(messageID uint) (MessageAssociateUser, error) {
	var messageAssociateUser MessageAssociateUser

	err := repository.db.
		Select("message.*, user.name as user_name").
		Joins("JOIN user on message.user_id = user.id").
		First(&messageAssociateUser, messageID).
		Error

	return messageAssociateUser, err
}

func (repository *MessageRepository) CreateMessage(message *Message) error {
	return repository.db.Create(&message).Error
}

func (repository *MessageRepository) UpdateMessage(message *Message) error {
	return repository.db.Model(&message).Updates(message).Error
}

func (repository *MessageRepository) DeleteMessage(message *Message) error {
	return repository.db.Delete(&message).Error
}
