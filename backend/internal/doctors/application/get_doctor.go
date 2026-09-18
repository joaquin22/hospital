package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	userApp "github.com/joaquin22/hospital-api/internal/users/application"
)

type GetDoctorUseCase struct {
	repo   domain.DoctorRepository
	userUC *userApp.GetUserUseCase
}

func NewGetDoctorUseCase(repo domain.DoctorRepository, userUC *userApp.GetUserUseCase) *GetDoctorUseCase {
	return &GetDoctorUseCase{repo: repo, userUC: userUC}
}

func (uc *GetDoctorUseCase) Execute(id uint) (*DoctorUserOutput, error) {
	doctor, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user, err := uc.userUC.Execute(doctor.UserID())
	if err != nil {
		return nil, err
	}

	return toDoctorUserOutput(doctor, user), nil
}
