package user

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	GetAll() ([]User, error)
	GetUserByID(userID uint) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repository *Repository) GetAll() ([]User, error) {
	var users []User
	err := repository.db.Find(&users).Error
	return users, err
}

func (repository *Repository) GetUserByID(userID uint) (User, error) {
	var user User
	err := repository.db.First(&user, userID).Error
	return user, err
}

func (repository *Repository) CreateUser(user *User) error {
	return repository.db.Create(&user).Error
}

func (repository *Repository) UpdateUser(user *User) error {
	return repository.db.Model(&user).Updates(user).Error
}

func (repository *Repository) DeleteUser(user *User) error {
	return repository.db.Delete(&user).Error
}
