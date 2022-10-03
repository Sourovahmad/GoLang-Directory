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
	router.POST("/category-delete", controllers.DeleteCategory)

	// books

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/book-create", controllers.CreateBook)

	// Default router run
	router.Run(":8080")

}
