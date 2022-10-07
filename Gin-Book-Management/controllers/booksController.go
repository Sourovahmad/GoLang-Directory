package controllers

import (
	"fmt"
	"gin-book-management/database"
	"gin-book-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {

	var books []models.Books
	var category models.Categories

	

	database.Database.Find(&books)

	for i := range books {
		currentBook := books[i]
		category_id := currentBook.Category_id
		database.Database.Where("id = ?", category_id).First(&category)

		fmt.Println(category)

	}

	allTheBooks := books

	c.IndentedJSON(http.StatusOK, gin.H{
		"books": allTheBooks,
	})
}

func CreateBook(c *gin.Context) {

	var request models.Books

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	}

	database.Database.Create(&request)

	c.IndentedJSON(http.StatusOK, gin.H{
		"new_book": request,
	})

}

func GetSingleBook(c *gin.Context) {
	//
}
