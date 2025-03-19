package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBrand(c *gin.Context) {
	var brand models.Brands
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	brand.IsDeleted = true
	if err := config.DB.Save(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}
