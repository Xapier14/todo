package todos

import (
	"github.com/gin-gonic/gin"

	"github.com/xapier14/todo/internal/controllers/todos"
	"github.com/xapier14/todo/internal/responses/general"
	responses "github.com/xapier14/todo/internal/responses/todos"
)

type PostTodoRequest struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
}

func PostTodo(c *gin.Context) {
	userId := c.GetUint("user_id")

	var request PostTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, general.GenerateBadRequestResponse("Invalid request body"))
		return
	}

	todo, err := todos.CreateTodo(userId, request.Title, request.Description, request.IsCompleted)
	if err != nil {
		c.AbortWithStatusJSON(500, general.GenerateInternalServerErrorResponse())
		return
	}

	c.JSON(201, responses.GenerateCreatedTodoResponse(todo))
}