package application

import (
	"github.com/joaquin22/hospital-api/internal/users/domain"
)

type RegisterUserUseCase struct {
	repo   domain.UserRepository
	hasher domain.PasswordHasher
}

func NewRegisterUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *RegisterUserUseCase) Execute(userInput RegisterUserInput) (*UsersOutput, error) {

	email, err := domain.NewEmail(userInput.Email)

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

	role, err := domain.NewRole(userInput.Role)
	if err != nil {
		return nil, err
	}

	dni, err := domain.NewDni(userInput.Dni)
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

	hashedPassword, err := uc.hasher.Hash(userInput.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(userInput.FullName, email, hashedPassword, dni, role)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}

	return toOutput(user), nil
}
