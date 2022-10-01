package models

type Books struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Category_id string `json:"category_id" validate:"required"`
	Author_id   string `json:"author_id" validate:"required"`
}
