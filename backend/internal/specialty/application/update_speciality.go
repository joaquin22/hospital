package application

import "github.com/joaquin22/hospital-api/internal/specialty/domain"

type UpdateSpecialityUseCase struct {
	repo domain.SpecialityRepository
}

func NewUpdateSpecialityUseCase(repo domain.SpecialityRepository) *UpdateSpecialityUseCase {
	return &UpdateSpecialityUseCase{repo: repo}
}

func (uc *UpdateSpecialityUseCase) Execute(input UpdateSpecialityInput) (*SpecialityOutput, error) {
	speciality, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if speciality == nil {
		return nil, domain.ErrSpecialityNotFound
	}

	if err := speciality.UpdateSpeciality(input.Name, input.Description); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(speciality); err != nil {
		return nil, err
	}

	return toOutput(speciality), nil
}
