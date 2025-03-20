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
	var modelsInput []models.Models
	if err := c.ShouldBindJSON(&modelsInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
		return
	}

	modelCodeMap := make(map[string]bool)
	var newModels []models.Models

	for _, model := range modelsInput {
		if _, exists := modelCodeMap[model.ModelCode]; exists {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Duplicate model code in request"})
			return
		}
		modelCodeMap[model.ModelCode] = true

		if exists, err := helpers.CheckBrandExists(model.BrandID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to check brand"})
			return
		} else if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Brand does not exist"})
			return
		}

		var existingModel models.Models
		if err := config.DB.Where("model_code = ?", model.ModelCode).First(&existingModel).Error; err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}

		newModels = append(newModels, model)
	}

	if len(newModels) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All model codes already exist"})
		return
	}
	if err := config.DB.CreateInBatches(&newModels, 100).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create models"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": newModels})
}
