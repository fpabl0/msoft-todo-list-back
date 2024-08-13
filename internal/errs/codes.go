package errs

// Code represents the error code type
type Code string

// Error codes
const (
	CodeInvalidToken       Code = "INVALID_TOKEN"
	CodeInvalidTokenFormat Code = "INVALID_TOKEN_FORMAT"
	CodeAccessDenied       Code = "ACCESS_DENIED"
	CodeNotFound           Code = "NOT_FOUND"
	CodeInvalidCreds       Code = "INVALID_CREDENTIALS"
	CodeUnexpected         Code = "UNEXPECTED_ERROR"
)
