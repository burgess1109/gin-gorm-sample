package user

import (
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	GetAll() ([]User, error)
	GetUserByID(userID uint) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (repository *UserRepository) GetAll() ([]User, error) {
	var users []User
	err := repository.db.Find(&users).Error
	return users, err
}

func (repository *UserRepository) GetUserByID(userID uint) (User, error) {
	var user User
	err := repository.db.First(&user, userID).Error
	return user, err
}

func (repository *UserRepository) CreateUser(user *User) error {
	return repository.db.Create(&user).Error
}

func (repository *UserRepository) UpdateUser(user *User) error {
	return repository.db.Model(&user).Updates(user).Error
}

func (repository *UserRepository) DeleteUser(user *User) error {
	return repository.db.Delete(&user).Error
}
