package application

import "github.com/joaquin22/hospital-api/internal/patients/domain"

type PatchPatientUseCase struct {
	repo domain.PatientRepository
}

func NewPatchPatientUseCase(repo domain.PatientRepository) *PatchPatientUseCase {
	return &PatchPatientUseCase{
		repo: repo,
	}
}

func (uc *PatchPatientUseCase) Execute(input PatchPatientInput) (*PatientOutput, error) {
	patient, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, domain.ErrPatientNotFound
	}

	firstName := patient.FirstName()

	if input.FirstName != nil {
		firstName = *input.FirstName
	}

	lastName := patient.LastName()

	if input.LastName != nil {
		lastName = *input.LastName
	}

	phone := patient.Phone()

	if input.Phone != nil {
		newPhone, err := domain.NewPhone(*input.Phone)
		if err != nil {
			return nil, err
		}
		phone = newPhone
	}

	dni := patient.Dni()

	if input.Dni != nil {
		newDni, err := domain.NewDni(*input.Dni)
		if err != nil {
			return nil, err
		}

		if newDni.String() != patient.Dni().String() {
			existing, err := uc.repo.FindByDni(newDni)
			if err != nil {
				return nil, err
			}
			if existing != nil {
				return nil, domain.ErrDniAlreadyExists
			}
		}
		dni = newDni
	}

	email := patient.Email()

	if input.Email != nil {
		newEmail, err := domain.NewEmail(*input.Email)
		if err != nil {
			return nil, err
		}

		if newEmail.String() != patient.Email().String() {
			existing, err := uc.repo.FindByEmail(newEmail)
			if err != nil {
				return nil, err
			}
			if existing != nil {
				return nil, domain.ErrEmailAlreadyExists
			}
		}

		email = newEmail
	}

	if err := patient.UpdatePatient(firstName, lastName, dni, email, phone); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(patient); err != nil {
		return nil, err
	}

	return toOutput(patient), nil
}
