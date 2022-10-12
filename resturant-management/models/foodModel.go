package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image"`
	Food_id    string             `json:"food_id" validate:"reqired"`
	Menu_id    *string            `json:"menu_id" validate:"reqired"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
