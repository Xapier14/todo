package todos

import "github.com/xapier14/todo/internal/responses"

type TodoNotFoundResponse struct {
	responses.Response
	Id    string `json:"id"`
}

func GenerateTodoNotFoundResponse(id string) TodoNotFoundResponse {
	return TodoNotFoundResponse{
		Response: responses.Response{
			Message: "Todo not found",
			Error: "Todo not found",
		},
		Id:    id,
	}
}