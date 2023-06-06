package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-resto-management/controllers"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/api/menus", controller.GetMenus())
	router.GET("/api/menus/:menu_id", controller.GetMenu())
	router.POST("/api/menus", controller.CreateMenu())
	router.PATCH("/api/menus/:menu_id", controller.UpdateMenu())
}
