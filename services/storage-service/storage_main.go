package main

import (
	"fmt"
	"storage-service/config"
	"storage-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.StorageRoutes(r)
	fmt.Println("🚀 Server running on port 8083")
	r.Run(":8083")
}
