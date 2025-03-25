package controllers

import (
	"errors"
	"model-service/config"
	"model-service/helpers"
	"model-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateModels(c *gin.Context) {
	var inputModels []models.Models
	if err := c.ShouldBindJSON(&inputModels); err != nil {
		var singleModel models.Models
		if err := c.ShouldBindJSON(&singleModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Brand Not Exists"})
			return
		}
		inputModels = append(inputModels, singleModel)
	}
	var newModels []models.Models
	for _, model := range inputModels {
		var existingModel models.Models
		var brand helpers.Brand
		if model.ModelCode == "" || model.ModelName == "" || model.BrandID.String() == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Model code, name, and brand ID cannot be empty"})
			return
		}
		if err := config.DB.Where("model_code = ?", model.ModelCode).First(&existingModel).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "error": "Model code already exists"})
			return
		}
		if err := config.DB.Where("id = ?", model.BrandID).Find(&brand).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Brand not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}
		if model.BrandID.String() != brand.ID.String() {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Brand ID does not match"})
			return
		}
		newModels = append(newModels, model)
	}
	if len(newModels) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All model codes already exist"})
		return
	}
	if err := config.DB.Create(&newModels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create models"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": newModels})
}
