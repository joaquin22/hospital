package domain

import "errors"

var (
	ErrInvalidEmail       = errors.New("user: email inválido")
	ErrEmailAlreadyExists = errors.New("user: el email ya está registrado")
	ErrInvalidDni         = errors.New("user: DNI inválido")
	ErrDniAlreadyExists   = errors.New("user: el DNI ya está registrado")
	ErrPatientNotFound    = errors.New("patient: paciente no encontrado")
	ErrInvalidPhone       = errors.New("user: teléfono inválido")
)
