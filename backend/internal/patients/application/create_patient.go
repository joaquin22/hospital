package application

import "github.com/joaquin22/hospital-api/internal/patients/domain"

type CreatePatientUseCase struct {
	repo domain.PatientRepository
}

func NewCreatePatientUseCase(repo domain.PatientRepository) *CreatePatientUseCase {
	return &CreatePatientUseCase{
		repo: repo,
	}
}

func (uc *CreatePatientUseCase) Execute(input CreatePatientInput) (*PatientOutput, error) {
	email, err := domain.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	existingEmail, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if existingEmail != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	dni, err := domain.NewDni(input.Dni)
	if err != nil {
		return nil, err
	}

	existingDni, err := uc.repo.FindByDni(dni)
	if err != nil {
		return nil, err
	}

	if existingDni != nil {
		return nil, domain.ErrDniAlreadyExists
	}

	phone, err := domain.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	patient, err := domain.NewPatient(input.FirstName, input.LastName, dni, email, phone)
	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(patient)
	if err != nil {
		return nil, err
	}

	return toOutput(patient), nil
}
