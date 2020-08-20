package common

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
