package routes

import (
	controllers "defect-service/controllers/defect"
	defectChoiceControllers "defect-service/controllers/defect_choice"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefectRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://pre-owned-app.vercel.app", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
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
		defect.PATCH("/defect/delete", controllers.DeleteDefect)
	}
	defectChoice := r.Group("/defect-choices")
	{
		defectChoice.GET("/defect-choices", defectChoiceControllers.GetDefectChoices)
		defectChoice.GET("/defect-choice", defectChoiceControllers.GetDefectChoiceById)
		defectChoice.GET("/defect-choice/defects", defectChoiceControllers.GetDefectChoicesByDefectId)
		defectChoice.POST("/defect-choice/create", defectChoiceControllers.CreateDefectChoices)
		defectChoice.PUT("/defect-choice/update", defectChoiceControllers.UpdateDefectChoice)
		defectChoice.PUT("/defect-choice/delete", defectChoiceControllers.DeleteDefectChoice)
	}
}
