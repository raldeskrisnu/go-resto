package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	router.GET("/api/tables", controller.GetTables())
	router.GET("/api/table/:table_id", controller.GetTable())
	router.POST("/api/tables", controller.CreateTable())
	router.PATCH("/api/table/:table_id", controller.UpdateTable())
}
