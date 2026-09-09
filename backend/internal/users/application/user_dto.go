package application

import "time"

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
