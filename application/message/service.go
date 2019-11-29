package message

import (
	"gin-gorm-sample/application/domain"
	messagePort "gin-gorm-sample/application/message/port"
	userPort "gin-gorm-sample/application/user/port"
)

type Service struct {
	messageRepo messagePort.Repository
	userRepo    userPort.Repository
}

func NewService(messageRepo messagePort.Repository, userRepo userPort.Repository) *Service {
	return &Service{
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

func (s Service) Get(messageID uint) (domain.Message, error) {
	return s.messageRepo.Get(messageID)
}

func (s Service) GetAll(userID uint) ([]domain.Message, error) {
	return s.messageRepo.GetAll(userID)
}

func (s Service) Create(message *domain.Message) error {
	_, err := s.userRepo.Get(message.UserID)
	if err != nil {
		return err
	}

	return s.messageRepo.Create(message)
}

func (s Service) Update(message *domain.Message) error {
	if message.UserID != 0 {
		_, err := s.userRepo.Get(message.UserID)
		if err != nil {
			return err
		}
	}

	return s.messageRepo.Update(message)
}

func (s Service) Delete(id uint) error {
	return s.messageRepo.Delete(id)
}
