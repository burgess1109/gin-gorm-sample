package user

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type GetResponseBody struct {
	Data User `json:"data"`
}

type GetAllResponseBody struct {
	Data []User `json:"data"`
}
