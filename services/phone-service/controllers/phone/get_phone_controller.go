package controllers

import (
	"net/http"
	"phone-service/config"
	"phone-service/helpers"
	"phone-service/models"

	"github.com/gin-gonic/gin"
)

func GetPhones(c *gin.Context) {
	var phones []models.Phones
	if err := config.DB.Where("is_deleted = false").Order("phone_code desc").Find(&phones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch phones"})
		return
	}
	c.JSON(http.StatusOK, phones)
}
func GetPhone(c *gin.Context) {
	var phones models.Phones
	id := c.Query("id")
	if err := config.DB.Where("id = ?", id).Find(&phones).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, phones)
}

func GetStoragesByModelID(c *gin.Context) {
	modelID := c.Query("model_id")
	if modelID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "model_id is required"})
		return
	}

	var model helpers.Model
	if err := config.DB.Where("id = ? AND is_deleted = false", modelID).First(&model).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Model not found"})
		return
	}

	var storages []helpers.Storages
	err := config.DB.
		Joins("JOIN phones ON phones.storage_id = storages.id").
		Where("phones.model_id = ? AND phones.is_deleted = false AND storages.is_deleted = false", modelID).
		Group("storages.id, storages.storage_code, storages.storage_value, storages.is_deleted").
		Find(&storages).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch storages"})
		return
	}

	c.JSON(http.StatusOK, storages)
}
func GetViewBrands(c *gin.Context) {
	var viewBrands []models.ViewBrands
	if err := config.DB.Where("is_deleted = false").Find(&viewBrands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch view brands"})
		return
	}
	c.JSON(http.StatusOK, viewBrands)
}
func GetViewModelsByBrandID(c *gin.Context) {
	var viewModels []models.ViewModels
	id := c.Query("brand_id")
	if err := config.DB.Where("brand_id =? And is_deleted = false", id).Find(&viewModels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch view models"})
		return
	}
	c.JSON(http.StatusOK, viewModels)
}

func GetViewStoragesByModelID(c *gin.Context) {
	var viewPhones []models.ViewPhones
	modelID := c.Query("model_id")
	if err := config.DB.Where("model_id = ? AND is_deleted = false", modelID).Find(&viewPhones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch view storages"})
		return
	}
	c.JSON(http.StatusOK, viewPhones)
}
func GetViewPhoneWithDuctionsByPhoneId(c *gin.Context) {
	var viewPhoneWithDeductions []models.ViewPhoneWithDeductions
	phoneId := c.Query("phone_id")
	if err := config.DB.Where("phone_id = ? ", phoneId).Find(&viewPhoneWithDeductions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch view phone with deductions"})
		return
	}
	c.JSON(http.StatusOK, viewPhoneWithDeductions)
}

func GetViewPhone(c *gin.Context) {
	var viewPhones models.ViewPhones
	brandId := c.Query("brand_id")
	modelId := c.Query("model_id")
	storageId := c.Query("storage_id")

	if err := config.DB.Where("brand_id = ? AND model_id = ? AND storage_id = ?", brandId, modelId, storageId).Find(&viewPhones).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, viewPhones)
}

func GetViewPhones(c *gin.Context) {
	var viewPhones []models.ViewPhones
	if err := config.DB.Where("is_deleted = false").Find(&viewPhones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to fetch view phones"})
		return
	}
	c.JSON(http.StatusOK, viewPhones)
}
