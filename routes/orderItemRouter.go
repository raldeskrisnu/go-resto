package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(router *gin.Engine) {
	router.GET("/api/orderitems", controller.GetOrderItems())
	router.GET("/api/orderitem/:orderitem_id", controller.GetOrderItem())
	router.GET("/api/orderitem-order/:order_id", controller.GetOrderItemByOrder())
	router.POST("/api/orderitem", controller.CreateOrderItem())
	router.PATCH("/api/orderitem/:orderitem_id", controller.UpdateOrderItem())
}
