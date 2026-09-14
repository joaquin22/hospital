package domain

import "errors"

var (
	ErrDoctorNotFound      = errors.New("doctor: doctor no encontrado")
	ErrDoctorAlreadyExists = errors.New("doctor: el doctor ya existe")
)
