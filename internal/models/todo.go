package models

type Todo struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}