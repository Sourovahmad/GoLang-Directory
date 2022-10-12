package routes

import (
	"resturant-management/controller"

	"github.com/gin-gonic/gin"
)

func MenuRouter(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menus", controller.GetMenus)
	incomingRoutes.GET("/menus/:id", controller.GetMenu)
	incomingRoutes.POST("/menus", controller.CreateMenu)
	incomingRoutes.PATCH("/menu/:menu_id", controller.UpdateMenu)
}
