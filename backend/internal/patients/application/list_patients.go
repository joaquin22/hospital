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

func (uc *ListPatientsUseCase) Execute() ([]*CreatePatientOutput, error) {
	patients, err := uc.repo.FindAll()

	if err != nil {
		return nil, err
	}

	outputs := make([]*CreatePatientOutput, 0, len(patients))
	for _, p := range patients {
		outputs = append(outputs, &CreatePatientOutput{
			ID:        p.ID(),
			FirstName: p.FirstName(),
			LastName:  p.LastName(),
			Email:     p.Email().String(),
			Dni:       p.Dni().String(),
			Phone:     p.Phone().String(),
			Active:    p.IsActive(),
			CreatedAt: p.CreatedAt(),
			UpdatedAt: p.UpdatedAt(),
		})
	}
	return outputs, nil
}
