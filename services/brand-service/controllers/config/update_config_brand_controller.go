package controller

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToggleStatusConfigBrand(c *gin.Context) {
	var configBrand models.ConfigBrands
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&configBrand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ConfigBrand not found"})
		return
	}
	configBrand.IsDeleted = !configBrand.IsDeleted
	if err := config.DB.Save(&configBrand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle config brand status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ConfigBrand status toggled successfully"})
}
