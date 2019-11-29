package user

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

	user, err := r.userService.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, convertToGetResponseBody(user))
}

func convertToGetResponseBody(user domain.User) *GetResponseBody {
	responseBody := &GetResponseBody{Data: User{}}
	responseBody.Data = User(user)

	return responseBody
}
