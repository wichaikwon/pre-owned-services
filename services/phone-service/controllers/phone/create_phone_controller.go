package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"phone-service/config"
	"phone-service/helpers"
	"phone-service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePhone(c *gin.Context) {
	var phone models.Phones
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
		return
	}

	var existingPhone models.Phones
	if err := config.DB.Where("phone_code = ?", phone.PhoneCode).First(&existingPhone).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "Phone code already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
		return
	}

	if err := config.DB.Create(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create phone"})
		return
	}

	configBrands, err := helpers.FetchConfigBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch config brands"})
		return
	}

	filteredConfigBrands := []helpers.ConfigBrand{}
	for _, configBrand := range configBrands {
		if configBrand.BrandID == phone.BrandID && !configBrand.IsDeleted {
			filteredConfigBrands = append(filteredConfigBrands, configBrand)
		}
	}

	if len(filteredConfigBrands) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "No valid config brands available"})
		return
	}

	processedDefectChoices := make(map[string]bool)

	for _, configBrand := range filteredConfigBrands {
		fmt.Printf("Fetching defect choices for defect_id: %s\n", configBrand.DefectID)
		defectChoices, err := helpers.FetchDefectChoices(configBrand.DefectID)
		if err != nil {
			fmt.Printf("Error fetching defect choices for defect_id %s: %v\n", configBrand.DefectID, err)
			continue
		}

		if len(defectChoices) == 0 {
			fmt.Printf("No defect choices available for defect_id: %s\n", configBrand.DefectID)
			continue
		}

		for _, choice := range defectChoices {
			if processedDefectChoices[choice.ID.String()] {
				continue
			}

			priceDeduction := models.PriceDeductions{
				ID:             uuid.New(),
				PhoneID:        phone.ID,
				ConfigBrandID:  configBrand.ID,
				DefectChoiceID: choice.ID,
				Deduction:      0,
				IsDeleted:      false,
			}
			if err := config.DB.Create(&priceDeduction).Error; err != nil {
				fmt.Printf("Error creating price deduction for defect_choice_id %s: %v\n", choice.ID, err)
				continue
			}

			processedDefectChoices[choice.ID.String()] = true
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": phone})
}
