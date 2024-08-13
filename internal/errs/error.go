package errs

import "fmt"

// Error defines an application error
type Error struct {
	Code    Code
	Message string
}

// Error implements Error interface
func (e Error) Error() string {
	if e.Message == "" {
		return string(e.Code)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// New creates a new application error with the especified code and message.
func New(code Code, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

// NewWithError creates a new application error with the especified code using the error string as message.
func NewWithError(code Code, err error) *Error {
	return &Error{Code: code, Message: err.Error()}
}
