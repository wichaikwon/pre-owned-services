package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateModel(c *gin.Context) {
	var model models.Models
	id := c.Query("id")
	if err := config.DB.Where("id=? ", id).Find(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&model)
	c.JSON(http.StatusOK, gin.H{"message": "Model updated successfully"})
}
