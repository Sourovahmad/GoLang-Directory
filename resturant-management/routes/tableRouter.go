package routes


import(

	"github.com/gin-gonic/gin"
	"resturant-management/controller"

 )


 func TableRoutes(incomingRoutes *gin.Engine)  {

	incomingRoutes.GET("/tables", controller.GetTables)
	incomingRoutes.GET("/tables/:id", controller.GetTable)
	incomingRoutes.POST("/tables", controller.CreateTable)
	incomingRoutes.PATCH("/table/:table_id", controller.UpdateTable)
 }