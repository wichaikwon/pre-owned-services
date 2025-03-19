package controllers

import (
	"brand-service/config"
	"brand-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBrand(c *gin.Context) {
	var brand models.Brands
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&brand).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&brand)
	c.JSON(http.StatusOK, brand)
}
