package application

import (
	"github.com/joaquin22/hospital-api/internal/users/domain"
)

type GetUserUseCase struct {
	repo domain.UserRepository
}

func NewGetUserUseCase(repo domain.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (uc *GetUserUseCase) Execute(id uint) (*UsersOutput, error) {
	user, err := uc.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	return toOutput(user), nil
}
