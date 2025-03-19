package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func DeletePhone(c *gin.Context) {
	var phone models.Phones
	id := c.Query("id")

	if err := config.DB.Where("id=?", id).Find(&phone).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	phone.IsDeleted = true
	if err := config.DB.Save(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete phone"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Phone deleted successfully"})
}
