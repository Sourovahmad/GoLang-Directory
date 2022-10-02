package models

type Books struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Category_id string `json:"category_id" binding:"required"`
	Author_id   string `json:"author_id" binding:"required"`
}
