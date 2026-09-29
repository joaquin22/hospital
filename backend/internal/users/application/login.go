package application

import (
	"github.com/joaquin22/hospital-api/internal/users/domain"
)

type LoginUserUseCase struct {
	repo           domain.UserRepository
	hasher         domain.PasswordHasher
	tokenGenerator domain.TokenGenerator
}

func NewLoginUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher, tokenGenerator domain.TokenGenerator) *LoginUserUseCase {
	return &LoginUserUseCase{
		repo:           repo,
		hasher:         hasher,
		tokenGenerator: tokenGenerator,
	}
}

func (uc *LoginUserUseCase) Execute(login LoginInput) (*LoginOutput, error) {
	email, err := domain.NewEmail(login.Email)

	if err != nil {
		return nil, err
	}

	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrInvalidCredentials
	}

	if !uc.hasher.Compare(user.PasswordHash(), login.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	// Se comprueba DESPUÉS de la contraseña a propósito: si se hiciera antes,
	// un atacante podría distinguir qué emails existen y están desactivados
	// probándolos sin saber la clave. Con este orden, sin la contraseña
	// correcta la respuesta siempre es la misma.
	if !user.Active() {
		return nil, domain.ErrInactiveUser
	}

	token, expiresInSecond, err := uc.tokenGenerator.GenerateToken(user.ID(), user.Role())
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token:     token,
		ExpiresIn: expiresInSecond,
	}, nil

}
