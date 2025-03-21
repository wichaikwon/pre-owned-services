package controllers

import (
	"errors"
	"net/http"
	"phone-service/config"
	"phone-service/helpers"
	"phone-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdatePhone(c *gin.Context) {
	var phone models.Phones
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&phone).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&phone)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": phone})
}
func UpdatePhones(c *gin.Context) {
	var phones []models.Phones
	if err := c.ShouldBindJSON(&phones); err != nil {
		var phone models.Phones
		if err := c.ShouldBindJSON(&phone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		phones = append(phones, phone)
	}

	var updatedPhones []models.Phones
	for _, phone := range phones {
		var existingPhone models.Phones
		if phone.PhoneCode == "" || phone.PhoneName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone code and name cannot be empty"})
			return
		}
		if err := config.DB.Where("phone_code = ?", phone.PhoneCode).First(&existingPhone).Error; err != nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		brand, err := helpers.FindBrandByID(phone.BrandID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
			return
		}
		model, err := helpers.FindModelByID(phone.ModelID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
			return
		}
		storage, err := helpers.FindStorageByID(phone.StorageID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Storage not found"})
			return
		}
		if phone.BrandID != brand.ID || phone.ModelID != model.ID || phone.StorageID != storage.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone details do not match the provided brand, model, or storage"})
			return
		}
		if model.BrandID != brand.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Model does not belong to the specified brand"})
			return
		}

		updatedPhones = append(updatedPhones, phone)
	}
	if len(updatedPhones) == 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "All phone codes already exist"})
		return
	}
	if err := config.DB.Save(&updatedPhones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update phones"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": updatedPhones})
}
