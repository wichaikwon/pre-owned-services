package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteModel(c *gin.Context) {
	var model models.Models
	id := c.Query("id")

	if err := config.DB.Where("id=?", id).Find(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	model.IsDeleted = true
	if err := config.DB.Save(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete model"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Model deleted successfully"})
}
