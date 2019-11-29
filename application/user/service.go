package user

import (
	"gin-gorm-sample/application/domain"
	"gin-gorm-sample/application/user/port"
)

type Service struct {
	userRepo port.Repository
}

func NewService(userRepo port.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s Service) Get(userID uint) (domain.User, error) {
	return s.userRepo.Get(userID)
}

func (s Service) GetAll() ([]domain.User, error) {
	return s.userRepo.GetAll()
}

func (s Service) Create(user *domain.User) error {
	return s.userRepo.Create(user)
}

func (s Service) Update(user *domain.User) error {
	return s.userRepo.Update(user)
}

func (s Service) Delete(id uint) error {
	return s.userRepo.Delete(id)
}
