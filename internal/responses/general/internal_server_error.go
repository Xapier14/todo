package general

import "github.com/xapier14/todo/internal/responses"

type InternalServerErrorResponse struct {
	responses.Response
}

func GenerateInternalServerErrorResponse() InternalServerErrorResponse {
	return InternalServerErrorResponse{
		Response: responses.Response{
			Message: "Something went wrong",
			Error: "Internal server error",
		},
	}
}

func GenerateDetailedInternalServerErrorResponse(message string) InternalServerErrorResponse {
	return InternalServerErrorResponse{
		Response: responses.Response{
			Message: message,
			Error: "Internal server error",
		},
	}
}