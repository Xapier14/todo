package check

import "github.com/xapier14/todo/internal/responses"

type CheckResponse struct {
	responses.Response
	LoggedIn bool   `json:"logged_in"`
	State    string `json:"state"`
	TokenID  string `json:"token_id"`
}