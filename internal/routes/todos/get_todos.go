package todos

import (
	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/todos"
	responses "github.com/xapier14/todo/internal/responses/todos"
)

func GetTodos(c *gin.Context) {
	userId := c.GetUint("user_id")
	todos, _ := todos.GetTodosFromUserID(userId)
	c.JSON(200, responses.GenerateTodosFoundResponse(todos))
}