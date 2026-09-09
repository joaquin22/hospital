package application

import "github.com/joaquin22/hospital-api/internal/patients/domain"

type CreatePatienUseCase struct {
	repo domain.PatientRepository
}

func NewCreatePatientUseCase(repo domain.PatientRepository) *CreatePatienUseCase {
	return &CreatePatienUseCase{
		repo: repo,
	}
}
