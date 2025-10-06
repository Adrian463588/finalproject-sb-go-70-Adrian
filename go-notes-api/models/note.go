package models

type Note struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required,min=5"`
	IsFavorite bool   `json:"is_favorite"`
}