package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
)

type CreateDoctorUseCase struct {
	repo domain.DoctorRepository
}

func NewCreateDoctorUseCase(repo domain.DoctorRepository) *CreateDoctorUseCase {
	return &CreateDoctorUseCase{repo: repo}
}

func (uc *CreateDoctorUseCase) Execute(input CreateDoctorInput) (*DoctorOutput, error) {
	doctor, err := domain.NewDoctor(input.UserID, input.SpecialtyID, input.LicenseNumber)
	if err != nil {
		return nil, err
	}
	if err := uc.repo.Save(doctor); err != nil {
		return nil, err
	}

	return toOutput(doctor), nil

}
