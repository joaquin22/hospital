package infrastructure

import (
	"gorm.io/gorm"

	doctorApp "github.com/joaquin22/hospital-api/internal/doctors/application"
	doctorPersistence "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/persistence"

	userApp "github.com/joaquin22/hospital-api/internal/users/application"
	userDomain "github.com/joaquin22/hospital-api/internal/users/domain"
	userPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
)

type TransactionalUpdateDoctor struct {
	db     *gorm.DB
	hasher userDomain.PasswordHasher
}

func NewTransactionalUpdateDoctor(db *gorm.DB, hasher userDomain.PasswordHasher) *TransactionalUpdateDoctor {
	return &TransactionalUpdateDoctor{db: db, hasher: hasher}
}

func (t *TransactionalUpdateDoctor) Execute(input doctorApp.UpdateDoctorInput) (*doctorApp.DoctorUserOutput, error) {
	var output *doctorApp.DoctorUserOutput

	err := t.db.Transaction(func(tx *gorm.DB) error {
		userRepo := userPersistence.NewGormUserRepository(tx)
		doctorRepo := doctorPersistence.NewGormDoctorRepository(tx)

		getUserUC := userApp.NewGetUserUseCase(userRepo)
		updateUserUC := userApp.NewUpdateUserUseCase(userRepo, t.hasher)
		updateDoctorUC := doctorApp.NewUpdateDoctorUseCase(doctorRepo, getUserUC)

		doctor, err := doctorRepo.FindByID(input.ID)
		if err != nil {
			return err
		}

		if _, err := updateUserUC.Execute(userApp.UpdateUserInput{
			ID:        doctor.UserID(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  input.Password,
			Role:      input.Role,
			Dni:       input.Dni,
		}); err != nil {
			return err // dispara rollback — el UPDATE del user también se deshace
		}

		doctorOutput, err := updateDoctorUC.Execute(input)
		if err != nil {
			return err // dispara rollback — el UPDATE del user también se deshace
		}

		output = doctorOutput
		return nil // commit
	})

	return output, err
}
