package controllers

import (
	"math"
	"net/http"
	"phone-service/config"
	"phone-service/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetConfigPriceDeductions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	perPage := 10
	var phoneDeductions []models.PriceDeductions
	offset := (page - 1) * perPage
	if err := config.DB.Order("id desc").Offset(offset).Limit(perPage).Find(&phoneDeductions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch phone deductions"})
		return
	}
	var total int64
	config.DB.Model(&models.PriceDeductions{}).Count(&total)
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	c.JSON(http.StatusOK, gin.H{
		"data":        phoneDeductions,
		"page":        page,
		"per_page":    perPage,
		"total":       total,
		"total_pages": totalPages,
	})
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
func GetPriceDeductionByPhoneID(c *gin.Context) {
	var phoneDeductions []models.PriceDeductions
	phoneID := c.Query("id")
	if err := config.DB.Where("phone_id = ?", phoneID).Find(&phoneDeductions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, phoneDeductions)
}
