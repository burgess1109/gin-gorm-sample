package user

type postRequestBody struct {
	Name  string `json:"name" binding:"required,gt=1,lt=50"`
	Sex   string `json:"sex" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type patchRequestBody struct {
	ID    uint   `json:"-"`
	Name  string `json:"name" binding:"omitempty,gt=1,lt=50"`
	Sex   string `json:"sex"`
	Email string `json:"email" binding:"omitempty,email"`
}
