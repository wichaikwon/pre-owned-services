package routes

import (
	configControllers "phone-service/controllers/config"
	phoneControllers "phone-service/controllers/phone"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func PhoneRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	phone := r.Group("/phones")
	{
		phone.GET("/phones", phoneControllers.GetPhones)
		phone.GET("/phone", phoneControllers.GetPhone)
		phone.GET("/price-deductions", configControllers.GetConfigPriceDeductions)
		phone.GET("/price-deduction", configControllers.GetConfigPriceDeductionByID)
		phone.GET("/price-deductions/phone", configControllers.GetPriceDeductionByPhoneID)
		phone.GET("/view-brands", phoneControllers.GetViewBrands)
		phone.GET("/view-models", phoneControllers.GetViewModelsByBrandID)
		phone.GET("storages", phoneControllers.GetStoragesByModelID)
		phone.GET("/view-storages", phoneControllers.GetViewStoragesByModelID)
		phone.GET("/view-phones", phoneControllers.GetViewPhones)
		phone.GET("/view-phone-with-deductions", phoneControllers.GetViewPhoneWithDuctionsByPhoneId)
		phone.GET("/view-phone", phoneControllers.GetViewPhone)
		phone.POST("final-price", phoneControllers.FinalPrice)
		phone.POST("/phone/create", phoneControllers.CreatePhones)
		phone.PATCH("/price-deductions/update", configControllers.UpdateDeductions)
		phone.PATCH("/phone/delete", phoneControllers.DeletePhone)
		phone.PUT("/phone/update", phoneControllers.UpdatePhone)
	}
}
