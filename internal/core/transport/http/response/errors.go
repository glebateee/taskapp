package core_http_response

type ErrorResponse struct {
	Error   string `json:"error" example:"error description from application"`
	Message string `json:"message" example:"human-readable description of error"`
}
