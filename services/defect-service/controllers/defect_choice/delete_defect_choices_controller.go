package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDefectChoice(c *gin.Context) {
	var defectChoice models.DefectChoices
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&defectChoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DefectChoice not found"})
		return
	}
	defectChoice.IsDeleted = true
	if err := config.DB.Save(&defectChoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete defect choice"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DefectChoice deleted successfully"})
}
