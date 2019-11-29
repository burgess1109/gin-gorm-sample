package domain

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Sex       string
	Email     string
	UpdatedAt time.Time
	CreatedAt time.Time
}
