package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func GetPhones(c *gin.Context) {
	var phones []models.Phones
	if err := config.DB.Where("is_deleted = false").Order("phone_code desc").Find(&phones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch phones"})
		return
	}
	c.JSON(http.StatusOK, phones)
}
func GetPhone(c *gin.Context) {
	var phones models.Phones
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&phones).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, phones)
}
