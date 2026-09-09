package domain

import (
	"regexp"
)

var dniRegex = regexp.MustCompile(`^\d{8}$`)

type Dni struct {
	value string
}

func NewDni(value string) (Dni, error) {
	if !dniRegex.MatchString(value) {
		return Dni{}, ErrInvalidDni
	}
	return Dni{value: value}, nil
}

func (d Dni) String() string { return d.value }
