package domain

import "errors"

var (
	ErrInvalidEmail       = errors.New("user: email inválido")
	ErrEmailAlreadyExists = errors.New("user: el email ya está registrado")
	ErrInvalidDni         = errors.New("user: DNI inválido")
	ErrDniAlreadyExists   = errors.New("user: el DNI ya está registrado")
	ErrUserNotFound       = errors.New("user: usuario no encontrado")
	ErrInvalidPhone       = errors.New("user: teléfono inválido")
)
