package ms_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Errors interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type msErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e msErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e msErr) Message() string {
	return e.ErrMessage
}

func (e msErr) Status() int {
	return e.ErrStatus
}

func (e msErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewMsError(message string, status int, err string, causes []interface{}) Errors {
	return msErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewMsErrorFromBytes(bytes []byte) (Errors, error) {
	var apiErr msErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) Errors {
	return msErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) Errors {
	return msErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) Errors {
	return msErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) Errors {
	result := msErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}