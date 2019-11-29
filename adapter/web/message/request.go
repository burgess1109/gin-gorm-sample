package message

type postRequestBody struct {
	Context string `json:"context" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
}

type patchRequestBody struct {
	ID      uint   `json:"-"`
	Context string `json:"context"`
	UserID  uint   `json:"user_id"`
}
