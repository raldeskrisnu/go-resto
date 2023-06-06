package main

import (
	"go-resto-management/database"
	middleware "go-resto-management/middleware"
	myrouter "go-resto-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"os"

	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	myrouter.UserRoutes(router)
	router.Use(middleware.Authentication())

	myrouter.FoodRoutes(router)
	myrouter.MenuRoutes(router)
	myrouter.TableRoutes(router)
	myrouter.OrderRoutes(router)
	myrouter.OrderItemRoutes(router)
	myrouter.InvoiceRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
