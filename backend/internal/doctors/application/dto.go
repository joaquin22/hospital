package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	userApp "github.com/joaquin22/hospital-api/internal/users/application"
)

type CreateDoctorInput struct {
	UserID        uint
	SpecialtyID   uint
	LicenseNumber string
}
type RegisterDoctorInput struct {
	FirstName     string
	LastName      string
	Email         string
	Password      string
	Role          string
	Dni           string
	SpecialtyID   uint
	LicenseNumber string
}

type RegisterDoctorOutput struct {
	User   *userApp.UsersOutput `json:"user"`
	Doctor *DoctorOutput        `json:"doctor"`
}

type DoctorOutput struct {
	ID            uint   `json:"id"`
	SpecialityID  uint   `json:"speciality_id"`
	LicenseNumber string `json:"license_number"`
}

type DoctorUserOutput struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Dni           string `json:"dni"`
	SpecialityID  uint   `json:"speciality_id"`
	LicenseNumber string `json:"license_number"`
}

func toOutput(d *domain.Doctor) *DoctorOutput {
	return &DoctorOutput{
		ID:            d.ID(),
		SpecialityID:  d.SpecialityID(),
		LicenseNumber: d.LicenseNumber(),
	}
}
