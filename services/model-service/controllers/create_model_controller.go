package controllers

import (
	"encoding/json"
	"io"
	"model-service/config"
	"model-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckBrandExists(brandId string) (bool, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080/brands/brand?id="+brandId, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}

	return false, nil
}

func validateModel(c *gin.Context, model models.Models) bool {
	var existingModel models.Models
	if err := config.DB.Where("model_code = ?", model.ModelCode).First(&existingModel).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "modelCode already exists"})
		return false
	} else if err.Error() != "record not found" {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "database error: " + err.Error()})
		return false
	}
	return true
}

func validateBrand(c *gin.Context, brandId string) bool {
	exists, err := CheckBrandExists(brandId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "failed to check brand: " + err.Error()})
		return false
	}
	if !exists {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "brand does not exist"})
		return false
	}
	return true
}

func CreateModel(c *gin.Context) {
	var model models.Models
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "Invalid request data"})
		return
	}

	if !validateModel(c, model) || !validateBrand(c, model.BrandID.String()) {
		return
	}

	if err := config.DB.Create(&model).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "failed to create model: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": model})
}
