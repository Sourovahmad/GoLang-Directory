package models

type Books struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Category_id int    `json:"category_id" binding:"required"`
	Author_id   int    `json:"author_id" binding:"required"`
}
