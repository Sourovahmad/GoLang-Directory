package routes


import(

	"github.com/gin-gonic/gin"
	"resturant-management/controller"

 )


 func InvoiceRoutes(incomingRoutes *gin.Engine)  {

	incomingRoutes.GET("/invoices", controller.GetInvoices)
	incomingRoutes.GET("/invoices/:id", controller.GetInvoice)
	incomingRoutes.POST("/invoices", controller.CreateInvoice)
	incomingRoutes.PATCH("/invoice/:invoice_id", controller.UpdateInvoice)
 }