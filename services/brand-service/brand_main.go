package main

import (
	"brand-service/config"
	"brand-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.BrandRoutes(r)
	fmt.Println("🚀 Server running on port 8081")
	r.Run(":8081")
}
