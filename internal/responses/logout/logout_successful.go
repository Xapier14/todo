package logout

import "github.com/xapier14/todo/internal/responses"

type SuccessfulLogoutResponse struct {
	responses.Response
}

func GenerateSuccessfulLoginResponse() SuccessfulLogoutResponse {
	return SuccessfulLogoutResponse{
		Response: responses.Response{
			Message: "Login successful",
			Error: "",
		},
	}
}