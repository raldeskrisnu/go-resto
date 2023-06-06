package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/foods", controller.GetFoods())
	incomingRoutes.GET("/api/:food_id", controller.GetFood())
	incomingRoutes.POST("/api/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
