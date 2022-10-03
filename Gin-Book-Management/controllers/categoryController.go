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

	type requestForm struct {
		Id          int    `json:"id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	var request requestForm
	var category models.Categories

	// request validate and assing to variable
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// finding category and assigning to
	database.Database.First(&category, request.Id)

	category.Name = request.Name
	category.Description = request.Description

	database.Database.Save(&category)

	c.IndentedJSON(http.StatusOK, gin.H{
		"category": category,
	})

}

func DeleteCategory(c *gin.Context) {

	type requestForm struct {
		Id int `json:"id" binding:"required"`
	}

	var id requestForm
	var category models.Categories

	err := c.ShouldBindJSON(&id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// find category and assing to model

	database.Database.First(&category, id.Id)

	if category.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "category not found with this id",
		})

		return
	}

	database.Database.Delete(&category, id.Id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"category": "deleted successfully",
	})

}
