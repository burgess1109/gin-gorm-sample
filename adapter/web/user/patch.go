package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r *Router) patch(c *gin.Context) {
	var user patchRequestBody

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = uint(id)

	if err := r.userService.Update(convertPatchRequestBodyToDomainUser(user)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update user success",
	})
}

func convertPatchRequestBodyToDomainUser(PatchRequestBody patchRequestBody) *domain.User {
	user := domain.User{
		ID:    PatchRequestBody.ID,
		Name:  PatchRequestBody.Name,
		Sex:   PatchRequestBody.Sex,
		Email: PatchRequestBody.Email,
	}
	return &user
}
