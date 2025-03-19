package main

import (
	"auth-service/config"
	"auth-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.AuthRoutes(r)
	fmt.Println("🚀 Server running on port 8088")
	r.Run(":8088")
}
