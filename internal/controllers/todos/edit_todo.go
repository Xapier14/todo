package todos

import (
	"github.com/xapier14/todo/internal/db"
	"github.com/xapier14/todo/internal/models"
)

func EditTodo(todo *models.Todo) (*models.Todo, error) {
	database := db.GetDB()
	var todoCheck models.Todo
	err := database.Where("id = ? AND user_id = ?", todo.ID, todo.UserID).First(&todoCheck).Error
	if err != nil {
		return todo, err
	}
	err = database.Save(todo).Error
	return todo, err
}

func EditTodoFromID(id uint, userId uint, title string, description string, isCompleted bool) (models.Todo, error) {
	database := db.GetDB()
	var todo models.Todo
	err := database.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error
	if err != nil {
		return todo, err
	}
	todo.Title = title
	todo.Description = description
	todo.IsCompleted = isCompleted
	err = database.Save(&todo).Error
	return todo, err
}