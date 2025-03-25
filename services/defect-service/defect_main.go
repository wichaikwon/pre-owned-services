package main

import (
	"defect-service/config"
	"defect-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.DefectRoutes(r)
	fmt.Println("🚀 Server running on port 8084")
	r.Run(":8084")
}
