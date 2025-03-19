package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func UpdatePhone(c *gin.Context) {
	var phone models.Phones
	id := c.Query("id")
	brandId := c.Query("brand_id")
	modelId := c.Query("model_id")
	if err := config.DB.Where("id=? AND brand_id=? AND model_id=?", id, brandId, modelId).Find(&phone).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&phone)
	c.JSON(http.StatusOK, gin.H{"message": "Phone updated successfully"})
}
