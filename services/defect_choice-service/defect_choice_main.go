package main

import (
	"defect_choice-service/config"
	"defect_choice-service/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
	})
	routes.DefectChoiceRoutes(r)
	fmt.Println("🚀 Server running on port 8085")
	r.Run(":8085")

}
