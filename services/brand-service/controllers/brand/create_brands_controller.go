package controllers

import (
	"brand-service/config"
	"brand-service/helpers"
	"brand-service/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBrands(c *gin.Context) {
	var brands []models.Brands
	if err := c.ShouldBindJSON(&brands); err != nil {
		var singleBrand models.Brands
		if err := c.ShouldBindJSON(&singleBrand); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}
		brands = append(brands, singleBrand)
	}

	var newBrands []models.Brands
	for _, brand := range brands {
		var existingBrand models.Brands
		if brand.BrandCode == "" || brand.BrandName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Brand code and name cannot be empty"})
			return
		}
		if err := config.DB.Where("brand_code = ?", brand.BrandCode).First(&existingBrand).Error; err == nil {
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Database error"})
			return
		}
		newBrands = append(newBrands, brand)
	}
	if len(newBrands) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All brand codes already exist"})
		return
	}
	if err := config.DB.Create(&newBrands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create brands"})
		return
	}
	defects, err := helpers.FetchDefects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch defects"})
		return
	}
	if len(defects) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "No defects data available"})
		return
	}
	var configBrands []models.ConfigBrands

	for _, brand := range newBrands {
		for _, defect := range defects {
			if defect.ID == "" || defect.DefectName == "" {
				fmt.Println("Skipping invalid defect:", defect)
				continue
			}
			parsedUUID, err := uuid.Parse(defect.ID)
			if err != nil {
				fmt.Println("Skipping invalid UUID defect:", defect)
				continue
			}

			configBrands = append(configBrands, models.ConfigBrands{
				BrandID:    brand.ID,
				BrandCode:  brand.BrandCode,
				BrandName:  brand.BrandName,
				DefectID:   parsedUUID,
				DefectCode: defect.DefectCode,
				DefectName: defect.DefectName,
			})
		}
	}
	if len(configBrands) > 0 {
		if err := config.DB.Create(&configBrands).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create config brands"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"success":       true,
		"brands":        newBrands,
		"config_brands": configBrands,
	})
}
