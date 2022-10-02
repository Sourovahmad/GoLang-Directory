package controllers

import (
	"gin-book-management/database"
	"gin-book-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(request *gin.Context) {
	var categories []models.Categories
	database.Database.Find(&categories)

	request.IndentedJSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func CreateCategory(c *gin.Context) {
	var category models.Categories

	err := c.ShouldBind(&category)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	database.Database.Create(&category)

	c.IndentedJSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func UpdateCategory(c *gin.Context) {

	type requestID struct {
		ID string `json:"id"`
	}

	var requestid requestID

	err := c.ShouldBindJSON(&requestid)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"change": requestid,
	})

}
