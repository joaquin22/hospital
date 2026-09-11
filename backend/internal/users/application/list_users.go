package application

import "github.com/joaquin22/hospital-api/internal/users/domain"

type ListUsersUseCase struct {
	repo domain.UserRepository
}

func NewListUsersUseCase(repo domain.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		repo: repo,
	}
}

func (uc *ListUsersUseCase) Execute() ([]*UsersOutput, error) {
	users, err := uc.repo.FindAll()

	if err != nil {
		return nil, err
	}

	outputs := make([]*UsersOutput, 0, len(users))
	for _, u := range users {
		outputs = append(outputs, toOutput(u))
	}
	return outputs, nil
}
