package message

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetMessagesByUserID(userID uint) ([]AssociateUser, error)
	GetMessageByID(messageID uint) (Message, error)
	CreateMessage(message *Message) error
	UpdateMessage(message *Message) error
	DeleteMessage(message *Message) error
}

type AssociateUser struct {
	Message
	UserName string `json:"user_name"`
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repository *Repository) GetMessagesByUserID(userID uint) ([]AssociateUser, error) {
	var messageAssociateUsers []AssociateUser

	sqlDB := repository.db.
		Select("message.*, user.name as user_name").
		Joins("JOIN user ON message.user_id = user.id")

	if userID != 0 {
		sqlDB = sqlDB.Where("message.user_id = ?", userID)
	}

	err := sqlDB.Find(&messageAssociateUsers).Error

	return messageAssociateUsers, err
}

func (repository *Repository) GetMessageByID(messageID uint) (AssociateUser, error) {
	var messageAssociateUser AssociateUser

	err := repository.db.
		Select("message.*, user.name as user_name").
		Joins("JOIN user on message.user_id = user.id").
		First(&messageAssociateUser, messageID).
		Error

	return messageAssociateUser, err
}

func (repository *Repository) CreateMessage(message *Message) error {
	return repository.db.Create(&message).Error
}

func (repository *Repository) UpdateMessage(message *Message) error {
	return repository.db.Model(&message).Updates(message).Error
}

func (repository *Repository) DeleteMessage(message *Message) error {
	return repository.db.Delete(&message).Error
}
