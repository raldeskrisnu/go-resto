package controllers

import (
	"github.com/gin-gonic/gin"
	"go-resto-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
