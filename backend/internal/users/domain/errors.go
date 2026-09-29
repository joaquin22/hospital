package domain

import "errors"

var (
	ErrInvalidEmail           = errors.New("user: email inválido")
	ErrEmailAlreadyExists     = errors.New("user: el email ya está registrado")
	ErrEmailNotFound          = errors.New("user: email no encontrado")
	ErrInvalidDni             = errors.New("user: DNI inválido")
	ErrDniAlreadyExists       = errors.New("user: el DNI ya está registrado")
	ErrDniNotFound            = errors.New("user: el DNI no existe en RENIEC")
	ErrDniVerificationFailed  = errors.New("user: no se pudo verificar el DNI contra RENIEC")
	ErrInvalidName            = errors.New("user: nombre inválido")
	ErrInvalidRole            = errors.New("user: rol inválido")
	ErrWeakPassword           = errors.New("user: la contraseña debe tener al menos 8 caracteres")
	ErrEmailAlreadyRegistered = errors.New("user: ya existe un usuario con ese email")
	ErrUserNotFound           = errors.New("user: usuario no encontrado")
	ErrInvalidCredentials     = errors.New("user: email o contraseña incorrectos")
	ErrInactiveUser           = errors.New("user: la cuenta está desactivada")
)
