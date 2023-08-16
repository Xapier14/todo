package todos

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func GetTodoFromID(id uint, userId uint) (models.Todo, error) {
	database := db.GetDB()
	var todo models.Todo
	err := database.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error
	return todo, err
}