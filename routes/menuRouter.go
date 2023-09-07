package routes

import (
	controller "go-resto-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/api/menus", controller.GetMenus())
	router.GET("/api/menu/:menu_id", controller.GetMenu())
	router.POST("/api/menus", controller.CreateMenu())
	router.PATCH("/api/menu/:menu_id", controller.UpdateMenu())
}
