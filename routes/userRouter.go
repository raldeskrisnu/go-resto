package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/users", controller.GetUser())
	incomingRoutes.GET("/api/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/api/users/signup", controller.SignUp())
	incomingRoutes.POST("/api/login", controller.Login())
}
