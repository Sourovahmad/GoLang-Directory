package routes


import(

	"github.com/gin-gonic/gin"
	"resturant-management/controller"

 )


 func FoodRoutes(incomingRoutes *gin.Engine)  {

	incomingRoutes.GET("/foods", controller.GetFoods)
	incomingRoutes.GET("/foods/:id", controller.GetFood)
	incomingRoutes.POST("/foods", controller.CreateFood)
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood)
 }

