package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r Router) post(c *gin.Context) {
	var user postRequestBody

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := r.userService.Create(convertPostRequestBodyToDomainUser(user)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create user successful",
	})
}

func convertPostRequestBodyToDomainUser(PostRequestBody postRequestBody) *domain.User {
	user := domain.User{
		Name:  PostRequestBody.Name,
		Sex:   PostRequestBody.Sex,
		Email: PostRequestBody.Email,
	}
	return &user
}
