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

	database.Database.Find(&books)

	c.IndentedJSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func CreateBook(c *gin.Context) {

	var request models.Books
	var singleData models.Books

	err := c.ShouldBindJSON(&request)

	if err != nil {
		fmt.Println(err.Error())
	}

	database.Database.Where("name = ?", request.Name).First(&singleData)

	if singleData.ID != 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Book with This name Already Exist",
		})
		return
	}

	database.Database.Create(&request)

	c.IndentedJSON(http.StatusOK, gin.H{
		"new_book": request,
	})

}
