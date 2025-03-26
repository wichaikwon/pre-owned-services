package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateDefect(c *gin.Context) {
	var defect models.Defects
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&defect).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Defect not found"})
		return
	}
	if err := c.ShouldBindJSON(&defect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&defect)
	c.JSON(http.StatusOK, defect)
}
