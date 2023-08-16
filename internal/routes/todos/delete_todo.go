package todos

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xapier14/todo/internal/controllers/todos"
	responses "github.com/xapier14/todo/internal/responses/todos"
)

func DeleteTodo(c *gin.Context) {
	textId := c.Param("id")
	id, err := strconv.ParseUint(textId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(404, responses.GenerateTodoNotFoundResponse(textId))
		return
	}
	userId := c.GetUint("user_id")
	err = todos.DeleteTodoFromID(uint(id), userId)
	if err != nil {
		c.AbortWithStatusJSON(404, responses.GenerateTodoNotFoundResponse(textId))
		return
	}
	c.JSON(200, responses.GenerateTodoDeletedResponse(textId))
}