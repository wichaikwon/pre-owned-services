package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStorages(c *gin.Context) {
	var storages []models.Storages
	if err := config.DB.Where("is_deleted = false").Order("storage_code desc").Find(&storages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch storages"})
		return
	}
	c.JSON(http.StatusOK, storages)
}

func GetStorageById(c *gin.Context) {
	var storage models.Storages
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&storage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, storage)
}
