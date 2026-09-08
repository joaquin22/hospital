package domain

import "errors"

var (
	ErrInvalidEmail           = errors.New("user: email inválido")
	ErrInvalidName            = errors.New("user: nombre inválido")
	ErrInvalidRole            = errors.New("user: rol inválido")
	ErrWeakPassword           = errors.New("user: la contraseña debe tener al menos 8 caracteres")
	ErrEmailAlreadyRegistered = errors.New("user: ya existe un usuario con ese email")
	ErrUserNotFound           = errors.New("user: usuario no encontrado")
	ErrInvalidCredentials     = errors.New("user: email o contraseña incorrectos")
)
