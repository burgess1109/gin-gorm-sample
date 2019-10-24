package messages

import (
	"time"
)

type Message struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Context   string    `json:"context"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (Message) TableName() string {
	return "message"
}
