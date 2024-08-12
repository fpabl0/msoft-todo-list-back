package errs

import "fmt"

// Error defines an application error
type Error struct {
	Code    Code
	Message string
}

// WithMessage assigns a message to the error
func (e *Error) WithMessage(m string) *Error {
	e.Message = m
	return e
}

// String implements Error interface
func (e *Error) String() string {
	if e.Message == "" {
		return string(e.Code)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// New creates a new application error with the especified code.
func New(code Code) *Error {
	return &Error{Code: code}
}
