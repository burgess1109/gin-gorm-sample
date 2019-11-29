package user

import (
	"github.com/jinzhu/gorm"

	"gin-gorm-sample/application/domain"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) GetAll() ([]domain.User, error) {
	var users []User
	var domainUsers []domain.User

	err := repository.db.Find(&users).Error

	for _, user := range users {
		domainUsers = append(domainUsers, ConvertToDomainUser(user))
	}

	return domainUsers, err
}

func (repository *Repository) Get(userID uint) (domain.User, error) {
	var user User
	err := repository.db.First(&user, userID).Error

	return ConvertToDomainUser(user), err
}

func (repository *Repository) Create(domainUser *domain.User) error {
	user := User{
		Name:  domainUser.Name,
		Sex:   domainUser.Sex,
		Email: domainUser.Email,
	}
	return repository.db.Create(&user).Error
}

func (repository *Repository) Update(domainUser *domain.User) error {
	user := User{
		ID:    domainUser.ID,
		Name:  domainUser.Name,
		Sex:   domainUser.Sex,
		Email: domainUser.Email,
	}
	return repository.db.Model(&user).Updates(user).Error
}

func (repository *Repository) Delete(id uint) error {
	user := User{
		ID: id,
	}
	return repository.db.Delete(&user).Error
}

func ConvertToDomainUser(user User) domain.User {
	domainUser := domain.User(user)
	return domainUser
}
