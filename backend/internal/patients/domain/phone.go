package domain

import (
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

type Phone struct {
	number string
}

func NewPhone(number string) (Phone, error) {
	if !phoneRegex.MatchString(number) {
		return Phone{}, ErrInvalidPhone
	}
	return Phone{number: number}, nil
}

func (p *Phone) Number() string { return p.number }
