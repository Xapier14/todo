package logout

import "github.com/xapier14/todo/internal/responses"

type LogoutErrorResponse struct {
	responses.Response
}

func GenerateLogoutErrorResponse() LogoutErrorResponse {
	return LogoutErrorResponse{
		Response: responses.Response{
			Message: "Error occurred while logging out.",
			Error: "Error occurred while logging out.",
		},
	}
}