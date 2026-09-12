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

func (uc *GetPatientUseCase) Execute(id uint) (*PatientOutput, error) {
	patient, err := uc.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return toOutput(patient), nil
}
