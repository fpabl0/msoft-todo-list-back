package user

import "errors"

// User related errors
var (
	ErrAccessDenied    = errors.New("Acceso denegado para este recurso")
	ErrInvalidTokenFmt = errors.New("Formato de token de acceso inválido")
	ErrInvalidToken    = errors.New("Token de acceso inválido")
)
