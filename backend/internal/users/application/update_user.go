package application

import "github.com/joaquin22/hospital-api/internal/users/domain"

type UpdateUserUseCase struct {
	repo   domain.UserRepository
	hasher domain.PasswordHasher
}

func NewUpdateUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo, hasher: hasher}
}

func (uc *UpdateUserUseCase) Execute(input UpdateUserInput) (*UsersOutput, error) {
	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	email, err := domain.NewEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if email.String() != user.Email().String() {
		existing, err := uc.repo.FindByEmail(email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, domain.ErrEmailAlreadyExists
		}
	}

	role, err := domain.NewRole(input.Role)
	if err != nil {
		return nil, err
	}

	dni, err := domain.NewDni(input.Dni)
	if err != nil {
		return nil, err
	}

	if dni.String() != user.Dni().String() {
		existing, err := uc.repo.FindByDni(dni)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, domain.ErrDniAlreadyExists
		}
	}

	hashedPassword, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	if err := user.Update(input.FirstName, input.LastName, email, hashedPassword, dni, role); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}

	return toOutput(user), nil
}
