package login

import "github.com/xapier14/todo/internal/responses"

type SuccessfulLoginResponse struct {
	responses.Response
	SessionToken string `json:"sessionToken"`
}

func GenerateSuccessfulLoginResponse(sessionToken string) SuccessfulLoginResponse {
	return SuccessfulLoginResponse{
		Response: responses.Response{
			Message: "Login successful",
			Error: "",
		},
		SessionToken: sessionToken,
	}
}