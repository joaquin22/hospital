package application

import "github.com/joaquin22/hospital-api/internal/patients/domain"

type ListPatientsUseCase struct {
	repo domain.PatientRepository
}

func NewListPatientsUseCase(repo domain.PatientRepository) *ListPatientsUseCase {
	return &ListPatientsUseCase{
		repo: repo,
	}
}

func (uc *ListPatientsUseCase) Execute() ([]*PatientOutput, error) {
	patients, err := uc.repo.FindAll()

	if err != nil {
		return nil, err
	}

	outputs := make([]*PatientOutput, 0, len(patients))
	for _, p := range patients {
		outputs = append(outputs, toOutput(p))
	}

	return outputs, nil
}
