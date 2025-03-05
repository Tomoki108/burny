package io

type ErrorResponse struct {
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewErrResp(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}
