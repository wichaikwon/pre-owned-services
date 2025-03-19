package controllers

import (
	"errors"
	"net/http"
	"storage-service/config"
	"storage-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateStorage(c *gin.Context) {
	var storage models.Storages
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
		return
	}
	var existingStorage models.Storages
	if err := config.DB.Where("storage_code = ?", storage.StorageCode).First(&existingStorage).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "storageCode already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "database error"})
		return
	}
	if err := config.DB.Create(&storage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "failed to create storage"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": storage})
}
