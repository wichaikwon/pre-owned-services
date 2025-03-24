package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/models"
	"sort"

	"github.com/gin-gonic/gin"
)

func FinalPrice(c *gin.Context) {
	var phone models.Phones
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&phone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type RequestBody struct {
		PhoneID        string   `json:"phoneId"`
		DefectChoiceID []string `json:"defectChoiceId"`
	}

	var requestBodies []RequestBody
	if err := c.ShouldBindJSON(&requestBodies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var deductions []float64
	for _, requestBody := range requestBodies {
		for _, defectChoiceID := range requestBody.DefectChoiceID {
			var price models.PriceDeductions
			if err := config.DB.Where("phone_id = ? AND defect_choice_id = ?", requestBody.PhoneID, defectChoiceID).Find(&price).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			deductions = append(deductions, price.Deduction)
		}
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(deductions)))
	for _, deduction := range deductions {
		if (phone.Price - deduction) >= phone.MinPrice {
			phone.Price -= deduction
		} else {
			continue
		}
	}

	c.JSON(http.StatusOK, phone.Price)
}
