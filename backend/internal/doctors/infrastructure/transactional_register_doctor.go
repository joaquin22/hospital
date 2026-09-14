package infrastructure

import (
	"gorm.io/gorm"

	doctorApp "github.com/joaquin22/hospital-api/internal/doctors/application"
	doctorPersistence "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/persistence"

	userApp "github.com/joaquin22/hospital-api/internal/users/application"
	userDomain "github.com/joaquin22/hospital-api/internal/users/domain"
	userPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
)

type TransactionalRegisterDoctor struct {
	db     *gorm.DB
	hasher userDomain.PasswordHasher
}

func NewTransactionalRegisterDoctor(db *gorm.DB, hasher userDomain.PasswordHasher) *TransactionalRegisterDoctor {
	return &TransactionalRegisterDoctor{db: db, hasher: hasher}
}

func (t *TransactionalRegisterDoctor) Execute(input doctorApp.RegisterDoctorInput) (*doctorApp.RegisterDoctorOutput, error) {
	var output *doctorApp.RegisterDoctorOutput

	err := t.db.Transaction(func(tx *gorm.DB) error {
		userRepo := userPersistence.NewGormUserRepository(tx)
		doctorRepo := doctorPersistence.NewGormDoctorRepository(tx)

		registerUserUC := userApp.NewRegisterUserUseCase(userRepo, t.hasher)
		doctorUC := doctorApp.NewCreateDoctorUseCase(doctorRepo)

		userOutput, err := registerUserUC.Execute(userApp.RegisterUserInput{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  input.Password,
			Role:      input.Role,
			Dni:       input.Dni,
		})
		if err != nil {
			return err // dispara rollback — acá todavía no se escribió nada
		}

		doctorOutput, err := doctorUC.Execute(doctorApp.CreateDoctorInput{
			UserID:        userOutput.ID,
			SpecialtyID:   input.SpecialtyID,
			LicenseNumber: input.LicenseNumber,
		})

		if err != nil {
			return err // dispara rollback — el INSERT del user también se deshace
		}

		output = &doctorApp.RegisterDoctorOutput{User: userOutput, Doctor: doctorOutput}
		return nil // commit
	})

	return output, err
}
