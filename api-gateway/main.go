package main

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		log.Printf("Forwarding request to: %s%s", target, c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
func main() {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.GET("/users", reverseProxy("http://localhost:8088"))
	}

	brandsRoutes := r.Group("/brands")
	{
		brandsRoutes.GET("/brands", reverseProxy("http://localhost:8081"))
		brandsRoutes.GET("/brand", reverseProxy("http://localhost:8081"))
		brandsRoutes.GET("/config-brands", reverseProxy("http://localhost:8081"))
		brandsRoutes.GET("/config-brand", reverseProxy("http://localhost:8081"))

	}
	modelRoutes := r.Group("/models")
	{
		modelRoutes.GET("/models", reverseProxy("http://localhost:8082"))
		modelRoutes.GET("/model", reverseProxy("http://localhost:8082"))
	}
	storageRoutes := r.Group("/storages")
	{
		storageRoutes.GET("/storages", reverseProxy("http://localhost:8083"))
		storageRoutes.GET("/storage", reverseProxy("http://localhost:8083"))
	}
	phoneRoutes := r.Group("/phones")
	{
		phoneRoutes.GET("/phones", reverseProxy("http://localhost:8086"))
		phoneRoutes.GET("/phone", reverseProxy("http://localhost:8086"))
		phoneRoutes.GET("/price-deductions", reverseProxy("http://localhost:8086"))
		phoneRoutes.GET("/price-deduction", reverseProxy("http://localhost:8086"))
	}
	defectRoutes := r.Group("/defects")
	{
		defectRoutes.GET("/defects", reverseProxy("http://localhost:8084"))
		defectRoutes.GET("/defect", reverseProxy("http://localhost:8084"))
	}
	defectChoiceRoutes := r.Group("/defect-choices")
	{
		defectChoiceRoutes.GET("/defect-choices", reverseProxy("http://localhost:8085"))
		defectChoiceRoutes.GET("/defect-choice", reverseProxy("http://localhost:8085"))
		defectChoiceRoutes.GET("/defect-choice/defects", reverseProxy("http://localhost:8085"))
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Gateway",
		})
	})
	r.Run(":8080")
}
