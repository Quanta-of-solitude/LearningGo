package main

import "github.com/gin-gonic/gin"

func testHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hey there!",
	})
}
func main() {
	server := gin.Default()

	server.GET("/test", testHandler)
	server.Run(":8080")
}
