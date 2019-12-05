package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-gorm-sample/application/domain"
)

func (r *Router) patch(c *gin.Context) {
	var message patchRequestBody

	if err := c.ShouldBindJSON(&message); err != nil {
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
	message.ID = uint(id)

	if err := r.messageService.Update(convertPatchRequestBodyToDomainMessage(message)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update user successful",
	})
}

func convertPatchRequestBodyToDomainMessage(PatchRequestBody patchRequestBody) *domain.Message {
	message := domain.Message{
		ID:      PatchRequestBody.ID,
		Context: PatchRequestBody.Context,
		UserID:  PatchRequestBody.UserID,
	}
	return &message
}
