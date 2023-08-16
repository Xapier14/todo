package todos

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/todos"
	"github.com/xapier14/todo/internal/responses/general"
	responses "github.com/xapier14/todo/internal/responses/todos"
)

type PatchTodoRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
}

func PatchTodo(c *gin.Context) {
	textId := c.Param("id")
	userId := c.GetUint("user_id")

	id, err := strconv.ParseUint(textId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, responses.GenerateTodoNotFoundResponse(textId))
		return
	}

	var request PatchTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, general.GenerateBadRequestResponse("Invalid request body"))
		return
	}

	todo, err := todos.GetTodoFromID(uint(id), userId)
	if err != nil {
		c.AbortWithStatusJSON(400, responses.GenerateTodoNotFoundResponse(textId))
		return
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Description != "" {
		todo.Description = request.Description
	}
	todo.IsCompleted = request.IsCompleted

	_, err = todos.EditTodo(&todo)
	if err != nil {
		c.AbortWithStatusJSON(500, general.GenerateInternalServerErrorResponse())
		return
	}

	c.JSON(200, responses.GenerateUpdatedTodoResponse(todo))
}