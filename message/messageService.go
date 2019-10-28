package message

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service struct {
	messageRepo Repository
}

type ServiceInterface interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewService(messageRepo Repository) *Service {
	return &Service{messageRepo: messageRepo}
}

func (u *Service) Get(c *gin.Context) {
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

func (u *Service) GetByID(c *gin.Context) {
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

func (u *Service) Create(c *gin.Context) {
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

func (u *Service) Update(c *gin.Context) {
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

func (u *Service) Delete(c *gin.Context) {
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
