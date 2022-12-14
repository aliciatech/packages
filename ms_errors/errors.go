package ms_errors

import "net/http"

type RestErr struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"code"`
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "bad_request",
		StatusCode: http.StatusBadRequest,
	}
}

func NewNotFound(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "not_found",
		StatusCode: http.StatusBadRequest,
	}
}

func NewConflict(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "conflict",
		StatusCode: http.StatusConflict,
	}
}

func NewUnAuthorized(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "unauthorised",
		StatusCode: http.StatusUnauthorized,
	}
}

func NewForbidden(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "forbidden",
		StatusCode: http.StatusForbidden,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Error:      "internal_server_error",
		StatusCode: http.StatusInternalServerError,
	}
}
