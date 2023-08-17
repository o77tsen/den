package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/o77tsen/den/initializers"
	"github.com/o77tsen/den/models"
)

func GetAllMessages(c *gin.Context) {
	var messages []models.Message

	results := initializers.DB.Find(&messages)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": messages})
}

func GetMessage(c *gin.Context) {
	id := c.Param("id")

	var message []models.Message

	result := initializers.DB.First(&message, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "invalid", "message": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": message})
}

func CreateMessage(c *gin.Context) {
	var payload *models.CreateMessageReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newMessage := models.Message{
		Author:  payload.Author,
		Message: payload.Message,
	}

	result := initializers.DB.Create(&newMessage)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMessage})
}

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Message{}, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "invalid", "message": "Message not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
