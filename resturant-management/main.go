package main

import (
	"os"
	"resturant-management/database"
	"resturant-management/routes"
	"resturant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)
func main() {
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}


	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRouter(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoute(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}