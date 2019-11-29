package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r Router) getAll(c *gin.Context) {
	users, err := r.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, convertToGetAllResponseBody(users))
}

func convertToGetAllResponseBody(users []domain.User) *GetAllResponseBody {
	responseBody := &GetAllResponseBody{Data: []User{}}

	for _, user := range users {
		responseBody.Data = append(responseBody.Data, User(user))
	}

	return responseBody
}
