package controller

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToggleStatusConfigBrand(c *gin.Context) {
	var configBrands []models.ConfigBrands

	if err := c.ShouldBindJSON(&configBrands); err != nil {
		var configBrand models.ConfigBrands
		if err := c.ShouldBindJSON(&configBrand); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}

		if err := config.DB.Model(&models.ConfigBrands{}).
			Where("id = ?", configBrand.ID).
			Update("is_deleted", gorm.Expr("NOT is_deleted")).
			Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to update record"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Record toggled successfully"})
		return
	}

	tx := config.DB.Begin()
	for _, configBrand := range configBrands {
		if err := tx.Model(&models.ConfigBrands{}).
			Where("id = ?", configBrand.ID).
			Update("is_deleted", gorm.Expr("NOT is_deleted")).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to toggle records"})
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Records toggled successfully"})
}
