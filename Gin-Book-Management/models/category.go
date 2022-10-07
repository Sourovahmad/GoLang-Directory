package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
