package application

import (
	"github.com/joaquin22/hospital-api/internal/patients/domain"
)

type UpdatePatientUseCase struct {
	repo domain.PatientRepository
}

func NewUpdatePatientUseCase(repo domain.PatientRepository) *UpdatePatientUseCase {
	return &UpdatePatientUseCase{repo: repo}
}

func (uc *UpdatePatientUseCase) Execute(input UpdatePatientInput) (*PatientOutput, error) {

	patient, err := uc.repo.FindByID(input.ID)

	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, domain.ErrPatientNotFound
	}

	email, err := domain.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if email.String() != patient.Email().String() {
		existing, err := uc.repo.FindByEmail(email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, domain.ErrEmailAlreadyExists
		}
	}

	dni, err := domain.NewDni(input.Dni)
	if err != nil {
		return nil, err
	}

	if dni.String() != patient.Dni().String() {
		existing, err := uc.repo.FindByDni(dni)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, domain.ErrDniAlreadyExists
		}
	}

	phone, err := domain.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	if err := patient.UpdatePatient(input.FirstName, input.LastName, dni, email, phone); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(patient); err != nil {
		return nil, err
	}

	return toOutput(patient), nil
}
