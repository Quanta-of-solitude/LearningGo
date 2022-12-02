package main

import (
	"os"

	routes "golang-jwt-project/routes"

	"github.com/gin-gonic/gin"
)

var port string

func main() {

	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access Granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access Granted for api-2"})
	})
	router.Run(":" + port)
}

//if declaring the handler for api-1/2 outisde
//func api1Handler(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"success": "Access Granted for api-1",
//	})
//}
