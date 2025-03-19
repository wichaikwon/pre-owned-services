package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDefect(c *gin.Context) {
	var defect models.Defects
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&defect).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Defect not found"})
		return
	}
	defect.IsDeleted = true
	if err := config.DB.Save(&defect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete defect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Defect deleted successfully"})
}
