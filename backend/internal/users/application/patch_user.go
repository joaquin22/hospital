package application

import "github.com/joaquin22/hospital-api/internal/users/domain"

type PatchUserUseCase struct {
	repo   domain.UserRepository
	hasher domain.PasswordHasher
}

func NewPatchUserUseCase(repo domain.UserRepository, hasher domain.PasswordHasher) *PatchUserUseCase {
	return &PatchUserUseCase{repo: repo, hasher: hasher}
}

func (uc *PatchUserUseCase) Execute(input PatchUserInput) (*UsersOutput, error) {

	user, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	firstName := user.FirstName()

	if input.FirstName != nil {
		firstName = *input.FirstName
	}

	lastName := user.LastName()

	if input.LastName != nil {
		lastName = *input.LastName
	}

	email := user.Email()

	if input.Email != nil {
		newEmail, err := domain.NewEmail(*input.Email)

		if err != nil {
			return nil, err
		}

		if newEmail.String() != user.Email().String() {
			existing, err := uc.repo.FindByEmail(newEmail)
			if err != nil {
				return nil, err
			}
			if existing != nil {
				return nil, domain.ErrEmailAlreadyExists
			}
		}
		email = newEmail
	}

	hashedPassword := user.PasswordHash()

	if input.Password != nil {
		newHashedPassword, err := uc.hasher.Hash(*input.Password)
		if err != nil {
			return nil, err
		}
		hashedPassword = newHashedPassword
	}

	dni := user.Dni()

	if input.Dni != nil {
		newDni, err := domain.NewDni(*input.Dni)
		if err != nil {
			return nil, err
		}
		if newDni.String() != user.Dni().String() {
			existing, err := uc.repo.FindByDni(newDni)
			if err != nil {
				return nil, err
			}
			if existing != nil {
				return nil, domain.ErrDniAlreadyExists
			}
		}
		dni = newDni
	}

	role := user.Role()

	if input.Role != nil {
		newRole, err := domain.NewRole(*input.Role)
		if err != nil {
			return nil, err
		}
		role = newRole
	}

	if err := user.Update(firstName, lastName, email, hashedPassword, dni, role); err != nil {
		return nil, err
	}

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}

	return toOutput(user), nil

}
