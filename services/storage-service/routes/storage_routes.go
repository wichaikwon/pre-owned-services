package routes

import (
	"storage-service/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	storage := r.Group("/storages")
	{
		storage.GET("/storages", controllers.GetStorages)
		storage.GET("/storage", controllers.GetStorageById)
		storage.POST("/storage/create", controllers.CreateStorage)
		storage.PUT("/storage/update", controllers.UpdateStorage)
		storage.PUT("/storage/delete", controllers.DeleteStorage)
	}
}
