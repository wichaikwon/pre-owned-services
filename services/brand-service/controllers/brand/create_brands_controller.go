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

func CreateBrand(c *gin.Context) {
	var brand models.Brands
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "Invalid request data"})
		return
	}

	var existingBrand models.Brands
	if err := config.DB.Where("brand_code = ?", brand.BrandCode).First(&existingBrand).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "database error"})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "brandCode already exists"})
		return
	}

	if err := config.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "failed to create brand"})
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

	for _, defect := range defects {

		if defect.ID == "" || defect.DefectName == "" {
			fmt.Println("Skipping invalid defect:", defect)
			continue
		}

		parsedUUID, err := uuid.Parse(defect.ID)
		if err != nil {
			continue
		}

		configBrand := models.ConfigBrands{
			BrandID:    brand.ID,
			BrandCode:  brand.BrandCode,
			BrandName:  brand.BrandName,
			DefectID:   parsedUUID,
			DefectCode: defect.DefectCode,
			DefectName: defect.DefectName,
		}

		if err := config.DB.Create(&configBrand).Error; err != nil {
			fmt.Println("Error inserting configBrand:", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": brand})
}
