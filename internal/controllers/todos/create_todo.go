package todos

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func CreateTodo(userId uint, title string, description string, isCompleted bool) (models.Todo, error) {
	database := db.GetDB()
	todo := models.Todo{
		UserID:      userId,
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
	}
	err := database.Create(&todo).Error
	return todo, err
}