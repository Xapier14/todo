package general

import "github.com/xapier14/todo/internal/responses"

type RequestNotAuthorizedResponse struct {
	responses.Response
}

func GenerateRequestNotAuthorizedResponse() RequestNotAuthorizedResponse {
	return RequestNotAuthorizedResponse{
		Response: responses.Response{
			Message: "You are not authorized to perform this action",
			Error:   "Request not authorized",
		},
	}
}

func GenerateDetailedRequestNotAuthorizedResponse(message string) RequestNotAuthorizedResponse {
	return RequestNotAuthorizedResponse{
		Response: responses.Response{
			Message: message,
			Error:   "Request not authorized",
		},
	}
}
