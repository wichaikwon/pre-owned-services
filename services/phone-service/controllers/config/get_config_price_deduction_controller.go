package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func GetConfigPriceDeductions(c *gin.Context) {
	var phoneDeductions []models.PriceDeductions
	if err := config.DB.Order("id desc").Find(&phoneDeductions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch phone deductions"})
		return
	}
	c.JSON(http.StatusOK, phoneDeductions)
}

func GetConfigPriceDeductionByID(c *gin.Context) {
	var phoneDeductions models.PriceDeductions
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&phoneDeductions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, phoneDeductions)
}
