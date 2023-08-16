package general

import "github.com/xapier14/todo/internal/responses"

type BadRequestResponse struct {
	responses.Response
}

func GenerateBadRequestResponse(message string) BadRequestResponse {
	return BadRequestResponse{
		responses.Response{
			Message: message,
			Error:  "Bad Request",
		},
	}
}
