package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDefect(c *gin.Context) {
	var defect models.Defects
	if err := c.ShouldBindJSON(&defect); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "Invalid request data"})
		return
	}
	var existingDefect models.Defects
	if err := config.DB.Where("defect_code = ?", defect.DefectCode).First(&existingDefect).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "defectCode already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "database error"})
		return
	}
	if err := config.DB.Create(&defect).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "failed to create defect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": defect})
}
