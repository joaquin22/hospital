package application

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	userApp "github.com/joaquin22/hospital-api/internal/users/application"
)

type PatchDoctorUseCase struct {
	repo   domain.DoctorRepository
	userUC *userApp.GetUserUseCase
}

func NewPatchDoctorUseCase(repo domain.DoctorRepository, userUC *userApp.GetUserUseCase) *PatchDoctorUseCase {
	return &PatchDoctorUseCase{repo: repo, userUC: userUC}
}

func (uc *PatchDoctorUseCase) Execute(input PatchDoctorInput) (*DoctorUserOutput, error) {
	doctor, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	specialtyID := doctor.SpecialityID()
	if input.SpecialtyID != nil {
		specialtyID = *input.SpecialtyID
	}

	licenseNumber := doctor.LicenseNumber()
	if input.LicenseNumber != nil {
		licenseNumber = *input.LicenseNumber
	}

	if err := doctor.UpdateDoctor(specialtyID, licenseNumber); err != nil {
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
