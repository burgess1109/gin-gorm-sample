package user

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string `gorm:"type:varchar(50);not null"`
	Sex       string `gorm:"not null"`
	Email     string `gorm:"type:varchar(50);unique;not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
