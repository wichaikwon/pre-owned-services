package routes

import (
	"defect_choice-service/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefectChoiceRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	defectChoice := r.Group("/defect-choices")
	{

		defectChoice.GET("/defect-choices", controllers.GetDefectChoices)
		defectChoice.GET("/defect-choice", controllers.GetDefectChoiceById)
		defectChoice.GET("/defect-choice/defects", controllers.GetDefectChoicesByDefectId)
		defectChoice.POST("/defect-choice/create", controllers.CreateDefectChoice)
		defectChoice.PUT("/defect-choice/update", controllers.UpdateDefectChoice)
		defectChoice.PUT("/defect-choice/delete", controllers.DeleteDefectChoice)
	}
}
