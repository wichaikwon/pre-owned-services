package main

import (
	"fmt"
	"phone-service/config"
	"phone-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.PhoneRoutes(r)
	fmt.Println("🚀 Server running on port 8086")
	r.Run(":8086")
}
