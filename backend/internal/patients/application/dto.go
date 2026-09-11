package application

import (
	"time"

	"github.com/joaquin22/hospital-api/internal/patients/domain"
)

type CreatePatientInput struct {
	FirstName string
	LastName  string
	Dni       string
	Email     string
	Phone     string
}

type PatientOutput struct {
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

type UpdatePatientInput struct {
	ID        uint
	FirstName string
	LastName  string
	Dni       string
	Email     string
	Phone     string
}

type PatchPatientInput struct {
	ID        uint
	FirstName *string
	LastName  *string
	Dni       *string
	Email     *string
	Phone     *string
}

func toOutput(patient *domain.Patient) *PatientOutput {
	return &PatientOutput{
		ID:        patient.ID(),
		FirstName: patient.FirstName(),
		LastName:  patient.LastName(),
		Dni:       patient.Dni().String(),
		Email:     patient.Email().String(),
		Phone:     patient.Phone().String(),
		Active:    patient.IsActive(),
		CreatedAt: patient.CreatedAt(),
		UpdatedAt: patient.UpdatedAt(),
	}
}
