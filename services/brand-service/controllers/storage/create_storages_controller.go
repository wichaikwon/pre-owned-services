package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStorage(c *gin.Context) {
	var inputStorages []models.Storages
	if err := c.ShouldBindJSON(&inputStorages); err != nil {
		var singleStorage models.Storages
		if err := c.ShouldBindJSON(&singleStorage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Storage Not Exists"})
			return
		}
		inputStorages = append(inputStorages, singleStorage)
	}
	var newStorages []models.Storages
	for _, storage := range inputStorages {
		var existingStorage models.Storages
		if storage.StorageCode == "" || storage.StorageValue == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Storage code and name cannot be empty"})
			return
		}
		if err := config.DB.Where("storage_code = ?", storage.StorageCode).First(&existingStorage).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "error": "Storage code already exists"})
			return
		}
		newStorages = append(newStorages, storage)
	}
	if len(newStorages) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All storage codes already exist"})
		return
	}
	if err := config.DB.Create(&newStorages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create storages"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": newStorages})
}
