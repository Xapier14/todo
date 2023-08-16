package login

import "github.com/xapier14/todo/internal/responses"

type SuccessfulRefreshResponse struct {
	responses.Response
	SessionToken string `json:"sessionToken"`
}

func GenerateSuccessfulRefreshResponse(sessionToken string) SuccessfulRefreshResponse {
	return SuccessfulRefreshResponse{
		Response: responses.Response{
			Message: "Successfully refreshed session",
			Error: "",
		},
		SessionToken: sessionToken,
	}
}