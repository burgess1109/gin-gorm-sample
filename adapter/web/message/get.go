package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r Router) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message, err := r.messageService.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, convertToGetResponseBody(message))
}

func convertToGetResponseBody(message domain.Message) *GetResponseBody {
	responseBody := &GetResponseBody{Data: Message{}}

	responseUser := User{
		ID:   message.User.ID,
		Name: message.User.Name,
	}

	responseMessage := Message{
		ID:        message.ID,
		User:      responseUser,
		Context:   message.Context,
		UpdatedAt: message.UpdatedAt,
		CreatedAt: message.CreatedAt,
	}

	responseBody.Data = responseMessage

	return responseBody
}
