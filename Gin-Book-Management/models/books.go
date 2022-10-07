package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Category_id int    `json:"category_id" binding:"required"`
}
