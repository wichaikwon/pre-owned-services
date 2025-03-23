package controllers

import (
	"model-service/config"
	"model-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModels(c *gin.Context) {
	var models []models.Models
	if err := config.DB.Where("is_deleted = false").Order("model_code desc").Find(&models).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch models"})
		return
	}
	c.JSON(http.StatusOK, models)
}

func GetModelById(c *gin.Context) {
	var models models.Models
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&models).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, models)
}

func GetModelByBrandId(c *gin.Context) {
	var models []models.Models
	brandId := c.Query("brand_id")
	if err := config.DB.Where("brand_id = ? AND is_deleted = false", brandId).Order("model_code asc").Find(&models).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch models"})
		return
	}
	c.JSON(http.StatusOK, models)
}
