package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDefectChoices(c *gin.Context) {
	var inputDefectChoices []models.DefectChoices
	if err := c.ShouldBindJSON(&inputDefectChoices); err != nil {
		var singleDefectChoice models.DefectChoices
		if err := c.ShouldBindJSON(&singleDefectChoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Defect Not Exists"})
			return
		}
		inputDefectChoices = append(inputDefectChoices, singleDefectChoice)
	}
	var newDefectChoices []models.DefectChoices
	for _, defectChoice := range inputDefectChoices {
		var existingDefectChoice models.DefectChoices
		var defect models.Defects
		if defectChoice.ChoiceCode == "" || defectChoice.ChoiceName == "" || defectChoice.DefectID.String() == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Choice code, name, and defect ID cannot be empty"})
			return
		}
		if err := config.DB.Where("choice_code = ?", defectChoice.ChoiceCode).First(&existingDefectChoice).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "error": "Choice code already exists"})
			return
		}
		if err := config.DB.Where("id = ?", defectChoice.DefectID).Find(&defect).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Defect not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}
		if defectChoice.DefectID.String() != defect.ID.String() {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Defect ID does not match"})
			return
		}
		newDefectChoices = append(newDefectChoices, defectChoice)
	}
	if len(newDefectChoices) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All choice codes already exist"})
		return
	}
	if err := config.DB.Create(&newDefectChoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create defect choices"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": newDefectChoices})
}
