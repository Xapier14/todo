package todos

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/todos"
	responses "github.com/xapier14/todo/internal/responses/todos"
)

// GetTodo godoc
// @Summary Get a todo
// @Description Get a todo by ID
// @Tags todos
// @Accept none
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo

// @Router /todos/{id} [get]
func GetTodo(c *gin.Context) {
	textId := c.Param("id")
	id, err := strconv.ParseUint(textId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(404, responses.GenerateTodoNotFoundResponse(textId))
		return
	}
	userId := c.GetUint("user_id")
	todo, err := todos.GetTodoFromID(uint(id), userId)
	if err != nil {
		c.AbortWithStatusJSON(404, responses.GenerateTodoNotFoundResponse(textId))
		return
	}
	c.JSON(200, responses.GenerateTodoFoundResponse(todo))
}