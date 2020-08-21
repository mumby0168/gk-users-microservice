package common

import (
	"encoding/json"
	"io"
	"net/http"
)

type StandardError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Extra   interface{} `json:"extra"`
}

func (e StandardError) Error() string {
	return e.Message
}

func NewStandardError(code string, message string) StandardError {
	return StandardError{
		Code:    code,
		Message: message,
	}
}

func NewExtraStandardError(code string, message string, extra interface{}) StandardError {
	err := NewStandardError(code, message)
	err.Extra = extra
	return err
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message,
	}
}

func ErrorChecks(w http.ResponseWriter, response interface{}) (bool, error) {

	switch t := response.(type) {

	case StandardError:
		w.WriteHeader(http.StatusBadRequest)
		return true, WriteJson(w, t)
	case NotFoundError:
		w.WriteHeader(http.StatusNoContent)
		return true, nil
	}

	return false, nil
}

func WriteJson(w io.Writer, data interface{}) error {
	err := json.NewEncoder(w).Encode(data)
	return err
}
