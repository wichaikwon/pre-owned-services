package routes

import (
	brandControllers "brand-service/controllers/brand"
	configControllers "brand-service/controllers/config"
	modelControllers "brand-service/controllers/model"
	storageControllers "brand-service/controllers/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BrandRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://pre-owned-app.vercel.app"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	brand := r.Group("/brands")
	{
		brand.GET("/brands", brandControllers.GetBrands)
		brand.GET("/brand", brandControllers.GetBrandByID)
		brand.GET("/config-brands", configControllers.GetConfigBrands)
		brand.GET("/config-brand", configControllers.GetConfigBrandByID)
		brand.GET("/config-brands/brand", configControllers.GetConfigBrandByBrandID)
		brand.POST("/brands/create", brandControllers.CreateBrands)
		brand.PUT("/brand/update", brandControllers.UpdateBrand)
		brand.PUT("/brand/delete", brandControllers.DeleteBrand)
		brand.PATCH("/config-brands/update", configControllers.ToggleStatusConfigBrand)
	}
	model := r.Group("/models")
	{
		model.GET("/models", modelControllers.GetModels)
		model.GET("/model", modelControllers.GetModelById)
		model.GET("/models/brand", modelControllers.GetModelByBrandId)
		model.POST("/models/create", modelControllers.CreateModels)
		model.PUT("/model/update", modelControllers.UpdateModel)
		model.PUT("/model/delete", modelControllers.DeleteModel)
	}
	storage := r.Group("/storages")
	{
		storage.GET("/storages", storageControllers.GetStorages)
		storage.GET("/storage", storageControllers.GetStorageById)
		storage.POST("/storages/create", storageControllers.CreateStorage)
		storage.PUT("/storage/update", storageControllers.UpdateStorage)
		storage.PUT("/storage/delete", storageControllers.DeleteStorage)
	}
}
