package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStorage(c *gin.Context) {
	var storage models.Storages
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&storage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Storage not found"})
		return
	}
	storage.IsDeleted = true
	if err := config.DB.Save(&storage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete storage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storage deleted successfully"})
}
