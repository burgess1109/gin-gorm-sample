package message

import "time"

type Message struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Context   string    `json:"context"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetResponseBody struct {
	Data Message `json:"data"`
}

type GetAllResponseBody struct {
	Data []Message `json:"data"`
}
