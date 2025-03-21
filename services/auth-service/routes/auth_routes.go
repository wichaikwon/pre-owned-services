package routes

import (
	"auth-service/controllers"
	"auth-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.GET("/users", middlewares.AuthMiddleware(), controllers.GetUsers)
		auth.POST("/refresh_token", controllers.RefreshToken)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/register", controllers.Register)
	}
}
