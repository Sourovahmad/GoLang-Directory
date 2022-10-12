package routes

import (
	"resturant-management/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orders", controller.GetOrders)
	incomingRoutes.GET("/orders/:id", controller.GetOrder)
	incomingRoutes.POST("/orders", controller.CreateOrder)
	incomingRoutes.PATCH("/order/:order_id", controller.UpdateOrder)
}
