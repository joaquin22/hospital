package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	userApp "github.com/joaquin22/hospital-api/internal/users/application"
)

type UpdateDoctorUseCase struct {
	repo   domain.DoctorRepository
	userUC *userApp.GetUserUseCase
}

func NewUpdateDoctorUseCase(repo domain.DoctorRepository, userUC *userApp.GetUserUseCase) *UpdateDoctorUseCase {
	return &UpdateDoctorUseCase{repo: repo, userUC: userUC}
}

func (uc *UpdateDoctorUseCase) Execute(input UpdateDoctorInput) (*DoctorUserOutput, error) {
	doctor, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if err := doctor.UpdateDoctor(input.SpecialtyID, input.LicenseNumber); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(doctor); err != nil {
		return nil, err
	}

	user, err := uc.userUC.Execute(doctor.UserID())
	if err != nil {
		return nil, err
	}

	return toDoctorUserOutput(doctor, user), nil
}
