package routes

import (
	brandControllers "brand-service/controllers/brand"
	configControllers "brand-service/controllers/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BrandRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "DELETE", "OPTIONS"},
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
		brand.POST("/brand/create", brandControllers.CreateBrand)
		brand.PUT("/brand/update", brandControllers.UpdateBrand)
		brand.PUT("/brand/delete", brandControllers.DeleteBrand)
		brand.PUT("/config-brand/update", configControllers.ToggleStatusConfigBrand)
	}

}
