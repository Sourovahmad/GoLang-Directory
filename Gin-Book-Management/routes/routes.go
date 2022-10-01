package routes

import (
	"gin-book-management/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	router := gin.Default()

	router.GET("/categories", controllers.GetAllCategories)
	router.Run(":8080")
}
