package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r Router) getAll(c *gin.Context) {
	queryUserID := c.DefaultQuery("user_id", "0")

	userID, err := strconv.Atoi(queryUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	messages, err := r.messageService.GetAll(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, convertToGetAllResponseBody(messages))
}

func convertToGetAllResponseBody(messages []domain.Message) *GetAllResponseBody {
	responseBody := &GetAllResponseBody{Data: []Message{}}

	for _, message := range messages {
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

		responseBody.Data = append(responseBody.Data, responseMessage)
	}

	return responseBody
}
