package application

import "github.com/joaquin22/hospital-api/internal/patients/domain"

type GetPatientUseCase struct {
	repo domain.PatientRepository
}

func NewGetPatientUseCase(repo domain.PatientRepository) *GetPatientUseCase {
	return &GetPatientUseCase{
		repo: repo,
	}
}

func (uc *GetPatientUseCase) GetPatient(id uint) (*CreatePatientOutput, error) {
	patient, err := uc.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &CreatePatientOutput{
		ID:        patient.ID(),
		FirstName: patient.FirstName(),
		LastName:  patient.LastName(),
		Email:     patient.Email().String(),
		Dni:       patient.Dni().String(),
		Phone:     patient.Phone().String(),
		Active:    patient.IsActive(),
		CreatedAt: patient.CreatedAt(),
		UpdatedAt: patient.UpdatedAt(),
	}, nil
}
