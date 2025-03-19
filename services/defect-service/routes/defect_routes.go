package routes

import (
	"defect-service/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefectRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	defect := r.Group("/defects")
	{
		defect.GET("/defects", controllers.GetDefects)
		defect.GET("/defect", controllers.GetDefectByID)
		defect.POST("/defect/create", controllers.CreateDefect)
		defect.PUT("/defect/update", controllers.UpdateDefect)
		defect.PUT("/defect/delete", controllers.DeleteDefect)
	}
}
