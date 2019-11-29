package domain

import (
	"time"
)

type Message struct {
	ID        uint
	UserID    uint
	User      User
	Context   string
	UpdatedAt time.Time
	CreatedAt time.Time
}
