package application

import "github.com/joaquin22/hospital-api/internal/specialty/domain"

type CreateSpecilityUseCase struct {
	repo domain.SpecialityRepository
}

func NewCreateSpecilityUseCase(repo domain.SpecialityRepository) *CreateSpecilityUseCase {
	return &CreateSpecilityUseCase{repo: repo}
}

func (uc *CreateSpecilityUseCase) Execute(input CreateSpecialityInput) (*SpecialityOutput, error) {
	speciality, err := domain.NewSpeciality(0, input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	if err := uc.repo.Save(speciality); err != nil {
		return nil, err
	}

	return toOutput(speciality), nil
}
