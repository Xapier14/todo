package todos

import (
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses"
)

type UpdatedTodoResponse struct {
	responses.Response
	Data models.Todo `json:"data"`
}

func GenerateUpdatedTodoResponse(todo models.Todo) UpdatedTodoResponse {
	return UpdatedTodoResponse{
		Response: responses.Response{
			Message: "Todo updated successfully",
			Error: "",
		},
		Data: todo,
	}
}
