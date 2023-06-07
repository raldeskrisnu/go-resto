package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-resto-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validation = validator.New()

func GetFood() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetFoods() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

//func round(num float64) int {
//
//}
//
//func toFixed(num float64, precision int) float64 {
//
//}
