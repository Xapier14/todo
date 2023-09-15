package check

import "github.com/xapier14/todo/internal/responses"

func GenerateLoggedInResponse(tokenId string) CheckResponse {
	return CheckResponse{
		Response: responses.Response{
			Message: "Checked authentication status",
			Error: "",
		},
		LoggedIn: true,
		State: "Logged in.",
		TokenID: tokenId,
	}
}
