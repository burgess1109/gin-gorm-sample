package users

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Birthday  time.Time `json:"birthday"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "user"
}
