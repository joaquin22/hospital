package application

import (
	"fmt"

	"github.com/joaquin22/hospital-api/internal/users/domain"
)

type RegisterUserUseCase struct {
	repo        domain.UserRepository
	hasher      domain.PasswordHasher
	dniVerifier domain.DniVerifier
}

func NewRegisterUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher, dniVerifier domain.DniVerifier) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		repo:        repo,
		hasher:      hasher,
		dniVerifier: dniVerifier,
	}
}

func (uc *RegisterUserUseCase) Execute(userInput RegisterUserInput) (*UsersOutput, error) {

	// Se valida primero, antes que nada externo: es un chequeo local y gratis,
	// y así una contraseña débil no gasta una llamada a RENIEC.
	if err := domain.ValidatePassword(userInput.Password); err != nil {
		return nil, err
	}

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

	exists, err := uc.dniVerifier.Exists(dni.String())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrDniVerificationFailed, err)
	}

	if !exists {
		return nil, domain.ErrDniNotFound
	}

	hashedPassword, err := uc.hasher.Hash(userInput.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(userInput.FirstName, userInput.LastName, email, hashedPassword, dni, role)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}

	return toOutput(user), nil
}
