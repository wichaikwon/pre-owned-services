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

func CreatePhones(c *gin.Context) {
	var phones []models.Phones

	if err := c.ShouldBindJSON(&phones); err != nil {
		var singlePhone models.Phones
		if err := c.ShouldBindJSON(&singlePhone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}
		phones = append(phones, singlePhone)
	}
	var newPhones []models.Phones
	for _, phone := range phones {
		var existingPhone models.Phones
		if err := config.DB.Where("phone_code = ?", phone.PhoneCode).First(&existingPhone).Error; err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}

		newPhones = append(newPhones, phone)
	}

	if len(newPhones) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All phone codes already exist"})
		return
	}

	if err := config.DB.CreateInBatches(&newPhones, 100).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create phones"})
		return
	}

	for _, phone := range newPhones {
		configBrands, err := helpers.FetchConfigBrands()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": configBrands})
			return
		}
		filteredConfigBrands := []helpers.ConfigBrand{}
		for _, configBrand := range configBrands {
			if configBrand.BrandID == phone.BrandID && !configBrand.IsDeleted {
				filteredConfigBrands = append(filteredConfigBrands, configBrand)
			}
		}

		if len(filteredConfigBrands) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": fmt.Sprintf("No valid config brands available for phone_code: %s", phone.PhoneCode)})
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

			var priceDeductions []models.PriceDeductions
			for _, choice := range defectChoices {
				if processedDefectChoices[choice.ID.String()] {
					continue
				}

				priceDeductions = append(priceDeductions, models.PriceDeductions{
					ID:             uuid.New(),
					PhoneID:        phone.ID,
					ConfigBrandID:  configBrand.ID,
					DefectChoiceID: choice.ID,
					Deduction:      0,
					IsDeleted:      false,
				})

				processedDefectChoices[choice.ID.String()] = true
			}
			if len(priceDeductions) > 0 {
				if err := config.DB.CreateInBatches(&priceDeductions, 100).Error; err != nil {
					fmt.Printf("Error creating price deductions for phone_code %s: %v\n", phone.PhoneCode, err)
				}
			}
		}
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "phones": newPhones})
}
