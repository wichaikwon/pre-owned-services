package controller

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConfigBrands(c *gin.Context) {
	var configBrand []models.ConfigBrands
	if err := config.DB.Order("id desc").Find(&configBrand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch price adjustments"})
		return
	}
	c.JSON(http.StatusOK, configBrand)
}

func GetConfigBrandByID(c *gin.Context) {
	var configBrand models.ConfigBrands
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Find(&configBrand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ConfigBrand not found"})
		return
	}
	c.JSON(http.StatusOK, configBrand)
}
func GetConfigBrandByBrandID(c *gin.Context) {
	var configBrands []models.ConfigBrands
	brandID := c.Query("id")
	if err := config.DB.Where("brand_id = ? AND is_deleted = FALSE", brandID).Find(&configBrands).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ConfigBrands not found"})
		return
	}
	c.JSON(http.StatusOK, configBrands)
}
