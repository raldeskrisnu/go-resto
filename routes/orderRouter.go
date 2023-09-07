package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/api/orders", controller.GetOrders())
	router.GET("/api/orders/:order_id", controller.GetOrder())
	router.POST("/api/order", controller.CreateOrder())
	router.PATCH("/api/orders/:order_id", controller.UpdateOrder())
}
