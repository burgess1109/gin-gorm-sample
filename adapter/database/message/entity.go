package message

import (
	"time"

	"gin-gorm-sample/adapter/database/user"
)

type Message struct {
	ID        uint `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserID    uint `gorm:"not null"`
	User      user.User
	Context   string `gorm:"not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (Message) TableName() string {
	return "message"
}
