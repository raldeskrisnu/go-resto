package routes

import "github.com/gin-gonic/gin"
import controller "go-resto-management/controllers"

func TableRoutes(router *gin.Engine) {
	router.GET("/api/tables", controller.GetTables())
	router.GET("/api/tables/:tables_id", controller.GetTable())
	router.POST("/api/tables", controller.CreateTable())
	router.PATCH("/api/tables/:tables_id", controller.UpdateTable())
}
