package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	userApp "github.com/joaquin22/hospital-api/internal/users/application"
)

type ListDoctorUseCase struct {
	userUC *userApp.GetUserUseCase
	repo   domain.DoctorRepository
}

func NewListDoctorUseCase(repo domain.DoctorRepository, userUC *userApp.GetUserUseCase) *ListDoctorUseCase {
	return &ListDoctorUseCase{repo: repo, userUC: userUC}
}

func (uc *ListDoctorUseCase) Execute() ([]*DoctorUserOutput, error) {
	doctors, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]*DoctorUserOutput, len(doctors))
	for i, doctor := range doctors {

		user, err := uc.userUC.Execute(doctor.UserID())
		if err != nil {
			return nil, err
		}

		outputs[i] = &DoctorUserOutput{
			ID:            doctor.ID(),
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Email:         user.Email,
			Dni:           user.Dni,
			Role:          user.Role,
			LicenseNumber: doctor.LicenseNumber(),
			SpecialityID:  doctor.SpecialityID(),
		}
	}
	return outputs, nil
}
