package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDefects(c *gin.Context) {
	var defects []models.Defects
	if err := config.DB.Order("defect_code desc").Where("is_deleted = false").Find(&defects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch defects"})
		return
	}
	c.JSON(http.StatusOK, defects)
}
func GetDefectByID(c *gin.Context) {
	var defect models.Defects
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&defect).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, defect)
}
