package main

import (
	"fmt"
	"model-service/config"
	"model-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.ModelRoutes(r)
	fmt.Println("🚀 Server running on port 8082")
	r.Run(":8082")
}
