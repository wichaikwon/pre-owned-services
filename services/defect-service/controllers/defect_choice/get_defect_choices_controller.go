package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDefectChoices(c *gin.Context) {
	var defectChoices []models.DefectChoices
	if err := config.DB.Where("is_deleted = false").Order("choice_code desc").Find(&defectChoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch defect choices"})
		return
	}
	c.JSON(http.StatusOK, defectChoices)
}

func GetDefectChoiceById(c *gin.Context) {
	var defectChoices models.DefectChoices
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&defectChoices).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, defectChoices)
}
func GetDefectChoicesByDefectId(c *gin.Context) {
	var defectChoices []models.DefectChoices
	defectId := c.Query("id")
	if err := config.DB.Where("defect_id = ?", defectId).Find(&defectChoices).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, defectChoices)
}
