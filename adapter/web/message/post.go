package message

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r Router) post(c *gin.Context) {
	var message postRequestBody

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := r.messageService.Create(convertPostRequestBodyToDomainMessage(message)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create user successful",
	})
}

func convertPostRequestBodyToDomainMessage(PostRequestBody postRequestBody) *domain.Message {
	message := domain.Message{
		Context: PostRequestBody.Context,
		UserID:  PostRequestBody.UserID,
	}
	return &message
}
