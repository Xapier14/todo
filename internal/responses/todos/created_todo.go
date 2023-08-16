package todos

import (
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses"
)

type CreatedTodoResponse struct {
	responses.Response
	Data models.Todo `json:"data"`
}

func GenerateCreatedTodoResponse(todo models.Todo) CreatedTodoResponse {
	return CreatedTodoResponse{
		Response: responses.Response{
			Message: "Todo created successfully",
			Error: "",
		},
		Data: todo,
	}
}