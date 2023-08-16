package todos

import (
	"github.com/xapier14/todo/internal/models"
	"github.com/xapier14/todo/internal/responses"
)

type TodosFoundResponse struct {
	responses.Response
	Todos []models.Todo `json:"todos"`
}

func GenerateTodosFoundResponse(todos []models.Todo) TodosFoundResponse {
	return TodosFoundResponse{
		Response: responses.Response{
			Message: "Todos found",
			Error: "",
		},
		Todos: todos,
	}
}