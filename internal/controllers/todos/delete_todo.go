package todos

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func DeleteTodoFromID(id uint, userId uint) error {
	database := db.GetDB()
	err := database.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Todo{}).Error
	return err
}