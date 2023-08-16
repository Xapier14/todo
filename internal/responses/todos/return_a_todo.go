package todos

import (
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses"
)

type TodoFoundResponse struct {
	responses.Response
	Todo models.Todo `json:"todo"`
}

func GenerateTodoFoundResponse(todo models.Todo) TodoFoundResponse {
	return TodoFoundResponse{
		Response: responses.Response{
			Message: "Todo found",
			Error: "",
		},
		Todo: todo,
	}
}