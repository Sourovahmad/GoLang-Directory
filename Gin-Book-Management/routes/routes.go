package routes

import (
	"gin-book-management/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {

	router := gin.Default()

	// category section
	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/category-create", controllers.CreateCategory)
	router.POST("/category-edit", controllers.UpdateCategory)

	// Default router run
	router.Run(":8080")

}
