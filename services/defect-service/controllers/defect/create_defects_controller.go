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
	var defects []models.Defects
	if err := c.ShouldBindJSON(&defects); err != nil {
		var singleDefect models.Defects
		if err := c.ShouldBindJSON(&singleDefect); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}
		defects = append(defects, singleDefect)
	}
	var newDefects []models.Defects
	for _, defect := range defects {
		var existingDefect models.Defects
		if defect.DefectCode == "" || defect.DefectName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Defect code and name cannot be empty"})
			return
		}
		if err := config.DB.Where("defect_code = ?", defect.DefectCode).First(&existingDefect).Error; err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}
		newDefects = append(newDefects, defect)
	}
	if len(newDefects) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All defect codes already exist"})
		return
	}
	if err := config.DB.Create(&newDefects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create defects"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Defects created successfully"})
}
