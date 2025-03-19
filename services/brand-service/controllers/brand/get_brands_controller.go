package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBrands(c *gin.Context) {
	var brands []models.Brands
	if err := config.DB.Order("brand_code desc").Where("is_deleted = false").Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch brands"})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func GetBrandByID(c *gin.Context) {
	var brand models.Brands
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}
