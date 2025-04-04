package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func UpdateDeductions(c *gin.Context) {
	var deductions []models.PriceDeductions
	if err := c.ShouldBindJSON(&deductions); err != nil {
		var singleDeduction models.PriceDeductions
		if err := c.ShouldBindJSON(&singleDeduction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}
		deductions = append(deductions, singleDeduction)
	}

	for _, deduction := range deductions {
		var existingDeduction models.PriceDeductions
		if err := config.DB.Where("id = ?", deduction.ID).First(&existingDeduction).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Record not found for ID " + deduction.ID.String()})
			return
		}

		existingDeduction.Deduction = deduction.Deduction
		existingDeduction.IsDeleted = deduction.IsDeleted

		if err := config.DB.Save(&existingDeduction).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to update record for ID " + deduction.ID.String()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Deductions updated successfully"})
}
