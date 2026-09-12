package application

import "github.com/joaquin22/hospital-api/internal/specialty/domain"

type ListSpecialitiesUseCase struct {
	repo domain.SpecialityRepository
}

func NewListSpecialitiesUseCase(specialityRepo domain.SpecialityRepository) *ListSpecialitiesUseCase {
	return &ListSpecialitiesUseCase{repo: specialityRepo}
}

func (uc *ListSpecialitiesUseCase) Execute() ([]*SpecialityOutput, error) {
	specialities, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]*SpecialityOutput, 0, len(specialities))
	for _, speciality := range specialities {
		outputs = append(outputs, toOutput(speciality))
	}

	return outputs, nil
}
