package check

import "github.com/xapier14/todo/internal/responses"

func GenerateNotLoggedInResponse(state string) CheckResponse {
	return CheckResponse{
		Response: responses.Response{
			Message: "Checked authentication status",
			Error: "",
		},
		LoggedIn: false,
		State: state,
		TokenID: "",
	}
}
