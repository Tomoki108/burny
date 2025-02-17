package io

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	ErrorResponse
	Details []ErrorDetail `json:"details"`
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

func NewValidationErrResp(message string) ValidationErrorResponse {
	return ValidationErrorResponse{
		ErrorResponse: ErrorResponse{
			Message: message,
		},
	}
}
