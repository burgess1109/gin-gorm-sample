package messages

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageService struct {
	messageRepo MessageRepository
}

type MessageServiceInterface interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewMessageService(messageRepo MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

func (u *MessageService) Get(c *gin.Context) {
	userID, err := strconv.Atoi(c.DefaultQuery("user_id", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := u.messageRepo.GetMessagesByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (u *MessageService) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.messageRepo.GetMessageByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (u *MessageService) Create(c *gin.Context) {
	var message Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Field validation failed",
		})
		return
	}

	if err := u.messageRepo.CreateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create message success",
	})
}

func (u *MessageService) Update(c *gin.Context) {
	var message Message

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

	if err := u.messageRepo.UpdateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update message success",
	})
}

func (u *MessageService) Delete(c *gin.Context) {
	var message Message

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	message.ID = uint(id)

	if err := u.messageRepo.DeleteMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete message success",
	})
}
