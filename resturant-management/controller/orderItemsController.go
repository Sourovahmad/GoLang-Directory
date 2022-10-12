package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems(c *gin.Context) {
	//
}

func GetOrderItem(c *gin.Context) {
	//
}

func CreateOrderItem(c *gin.Context) {
	//
}

func UpdateOrderItem(c *gin.Context) {
	//
}

func GetOrderItemsByOrder(c *gin.Context) {
	//
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	//
}
