package login

import "github.com/xapier14/todo/internal/responses"

type InvalidCredentialsResponse struct {
	responses.Response
}

func GenerateInvalidCredentialsResponse() InvalidCredentialsResponse {
	return InvalidCredentialsResponse{
		Response: responses.Response{
			Message: "Invalid credentials",
			Error: "Invalid credentials",
		},
	}
}
