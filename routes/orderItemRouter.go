package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-resto-management/controllers"
)

func OrderItemRoutes(router *gin.Engine) {
	router.GET("/api/orderitems", controller.GetOrderItems())
	router.GET("/api/orderitems/:orderitem_id", controller.GetOrderItem())
	router.GET("/api/orderitems-order/:order_id", controller.GetOrderItemByOrder())
	router.POST("/api/orderitems", controller.CreateOrderItem())
	router.PATCH("/api/orderitems/:orderitem_id", controller.UpdateOrderItem())
}
