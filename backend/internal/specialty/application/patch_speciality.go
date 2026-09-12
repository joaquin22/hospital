package application

import "github.com/joaquin22/hospital-api/internal/specialty/domain"

type PatchSpecialityUseCase struct {
	repo domain.SpecialityRepository
}

func NewPatchSpecialityUseCase(repo domain.SpecialityRepository) *PatchSpecialityUseCase {
	return &PatchSpecialityUseCase{repo: repo}
}

func (uc *PatchSpecialityUseCase) Execute(input PatchSpecialityInput) (*SpecialityOutput, error) {
	speciality, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.Name != nil || input.Description != nil {
		name := speciality.Name()
		description := speciality.Description()
		if input.Name != nil {
			name = *input.Name
		}
		if input.Description != nil {
			description = *input.Description
		}
		if err := speciality.UpdateSpeciality(name, description); err != nil {
			return nil, err
		}
	}

	if err := uc.repo.Save(speciality); err != nil {
		return nil, err
	}

	return toOutput(speciality), nil
}
