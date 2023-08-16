package todos

import "github.com/xapier14/todo/internal/responses"

type TodoDeletedResponse struct {
	responses.Response
	Id      string `json:"id"`
}

func GenerateTodoDeletedResponse(id string) TodoDeletedResponse {
	return TodoDeletedResponse{
		Response: responses.Response{
			Message: "Todo deleted successfully",
			Error: "",
		},
		Id:      id,
	}
}