package application

import (
	"time"

	"github.com/joaquin22/hospital-api/internal/users/domain"
)

type RegisterUserInput struct {
	FullName string
	Email    string
	Password string
	Role     string
	Dni      string
}

type RegisterUserOutput struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Dni       string    `json:"dni"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

type UsersOutput struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Dni       string    `json:"dni"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toOutput(u *domain.User) *UsersOutput {
	return &UsersOutput{
		ID:        u.ID(),
		FullName:  u.FullName(),
		Email:     u.Email().String(),
		Role:      string(u.Role()),
		Dni:       u.Dni().String(),
		Active:    u.Active(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}
}
