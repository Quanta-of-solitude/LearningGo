package routes

import (
	controller "golang-jwt-project/controllers"
	middleware "golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	// we use middleware to maintain protected access
	//because after login the token is generated not before
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())
}
