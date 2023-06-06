package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"id"`
	Quantity      *string            `json:"quantity" validate:"required,eq=5|eq=M|eq=L"`
	Unit_Price    *float64           `json:"unit_price" validate:"required"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id" validate:"required"`
}
