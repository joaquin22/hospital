package application

import "github.com/joaquin22/hospital-api/internal/users/domain"

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

	token, expiresInSecond, err := uc.tokenGenerator.GenerateToken(user.ID(), user.Role())
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token:     token,
		ExpiresIn: expiresInSecond,
	}, nil

}
