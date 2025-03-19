package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func UpdateConfigPriceDeduction(c *gin.Context) {
	var phoneDeduction models.PriceDeductions
	id := c.Query("id")
	if err := config.DB.Where("id=?", id).Find(&phoneDeduction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	phoneDeduction.IsDeleted = !phoneDeduction.IsDeleted
	if err := config.DB.Save(&phoneDeduction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update phone deduction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phone deduction updated successfully"})
}
