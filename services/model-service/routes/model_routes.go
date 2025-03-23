package routes

import (
	"model-service/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ModelRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	model := r.Group("/models")
	{
		model.GET("/models", controllers.GetModels)
		model.GET("/model", controllers.GetModelById)
		model.GET("/models/brand", controllers.GetModelByBrandId)
		model.POST("/models/create", controllers.CreateModels)
		model.PUT("/model/update", controllers.UpdateModel)
		model.PUT("/model/delete", controllers.DeleteModel)
	}
}
