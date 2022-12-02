package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("users/signup", controller.signup()) // the handle fucntions for these routes are from the controller package
	incommingRoutes.POST("users/login", controller.login())
}
