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
		AllowMethods:     []string{"PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "application/json"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	phone := r.Group("/phones")
	{
		phone.GET("/phones", phoneControllers.GetPhones)
		phone.GET("/phone", phoneControllers.GetPhone)
		phone.POST("/phone/create", phoneControllers.CreatePhone)
		phone.PUT("/phone/update", phoneControllers.UpdatePhone)
		phone.PATCH("/phone/delete", phoneControllers.DeletePhone)
		phone.GET("/price-deductions", configControllers.GetConfigPriceDeductions)
		phone.GET("/price-deduction", configControllers.GetConfigPriceDeductionByID)
		phone.PATCH("/price-deduction/update", configControllers.UpdateConfigPriceDeduction)
	}
}
