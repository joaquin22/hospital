package application

import "github.com/joaquin22/hospital-api/internal/specialty/domain"

type GetSpecialityUseCase struct {
	repo domain.SpecialityRepository
}

func NewGetSpecialityUseCase(repo domain.SpecialityRepository) *GetSpecialityUseCase {
	return &GetSpecialityUseCase{repo: repo}
}

func (uc *GetSpecialityUseCase) Execute(id uint) (*SpecialityOutput, error) {
	speciality, err := uc.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return toOutput(speciality), nil
}
