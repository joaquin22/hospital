package infrastructure

import (
	"gorm.io/gorm"

	doctorApp "github.com/joaquin22/hospital-api/internal/doctors/application"
	doctorPersistence "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/persistence"

	userApp "github.com/joaquin22/hospital-api/internal/users/application"
	userDomain "github.com/joaquin22/hospital-api/internal/users/domain"
	userPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
)

type TransactionalPatchDoctor struct {
	db     *gorm.DB
	hasher userDomain.PasswordHasher
}

func NewTransactionalPatchDoctor(db *gorm.DB, hasher userDomain.PasswordHasher) *TransactionalPatchDoctor {
	return &TransactionalPatchDoctor{db: db, hasher: hasher}
}

func (t *TransactionalPatchDoctor) Execute(input doctorApp.PatchDoctorInput) (*doctorApp.DoctorUserOutput, error) {
	var output *doctorApp.DoctorUserOutput
	err := t.db.Transaction(func(tx *gorm.DB) error {
		userRepo := userPersistence.NewGormUserRepository(tx)
		doctorRepo := doctorPersistence.NewGormDoctorRepository(tx)

		getUserUC := userApp.NewGetUserUseCase(userRepo)
		patchUserUC := userApp.NewPatchUserUseCase(userRepo, t.hasher)
		patchDoctorUC := doctorApp.NewPatchDoctorUseCase(doctorRepo, getUserUC)

		doctor, err := doctorRepo.FindByID(input.ID)
		if err != nil {
			return err
		}

		if _, err := patchUserUC.Execute(userApp.PatchUserInput{
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
		doctorOutput, err := patchDoctorUC.Execute(input)
		if err != nil {
			return err // dispara rollback — el UPDATE del user también se deshace
		}

		output = doctorOutput
		return nil // commit
	})

	return output, err
}
