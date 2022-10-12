package routes

import (
	"resturant-management/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRoute(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems)
	incomingRoutes.GET("/orderItems/:id", controller.GetOrderItem)
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem)
	incomingRoutes.PATCH("/orderItem/:orderItem_id", controller.UpdateOrderItem)

	// Get order items by order id

	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder)

}
