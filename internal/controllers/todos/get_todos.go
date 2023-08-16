package todos

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func GetTodosFromUserID(userId uint) ([]models.Todo, error) {
	database := db.GetDB()
	var todos []models.Todo
	err := database.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}