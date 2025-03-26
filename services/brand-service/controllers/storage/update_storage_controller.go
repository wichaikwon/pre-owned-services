package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStorage(c *gin.Context) {
	var storage models.Storages
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&storage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Storage not found"})
		return
	}
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&storage)
	c.JSON(http.StatusOK, storage)
}
