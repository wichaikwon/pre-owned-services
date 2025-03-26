package controllers

import (
	"defect-service/config"
	"defect-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateDefectChoice(c *gin.Context) {
	var defectChoice models.DefectChoices
	id := c.Query("id")
	defectId := c.Query("defect_id")
	if err := config.DB.Where("id=? AND defect_id=?", id, defectId).Find(&defectChoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DefectChoice not found"})
		return
	}
	if err := c.ShouldBindJSON(&defectChoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&defectChoice)
	c.JSON(http.StatusOK, defectChoice)
}
