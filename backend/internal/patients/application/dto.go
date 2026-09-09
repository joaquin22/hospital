package application

import "time"

type CreatePatientInput struct {
	FirstName string
	LastName  string
	Dni       string
	Email     string
	Phone     string
}

type CreatePatientOutput struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Dni       string    `json:"dni"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
