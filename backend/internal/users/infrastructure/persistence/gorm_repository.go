package persistence

import (
	"errors"

	"github.com/joaquin22/hospital-api/internal/users/domain"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user *domain.User) error {
	model := &UserModel{
		ID:        user.ID(),
		Email:     user.Email().String(),
		Password:  user.PasswordHash(),
		Dni:       user.Dni().String(),
		Role:      string(user.Role()),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
	if err := r.db.Save(model).Error; err != nil {
		return err
	}
	user.SyncPersisted(model.ID, model.CreatedAt, model.UpdatedAt)

	return nil
}

func (r *GormUserRepository) FindByID(id string) (*domain.User, error) {
	var model UserModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user, err := toDomain(&model)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *GormUserRepository) FindByEmail(email domain.Email) (*domain.User, error) {
	var model UserModel
	if err := r.db.First(&model, "email = ?", email.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return toDomain(&model)
}

func (r *GormUserRepository) FindByDni(dni domain.Dni) (*domain.User, error) {
	var model UserModel
	if err := r.db.First(&model, "dni = ?", dni.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return toDomain(&model)
}

func (r *GormUserRepository) FindAll() ([]*domain.User, error) {
	var models []UserModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0, len(models))
	for _, model := range models {
		user, err := toDomain(&model)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func toDomain(m *UserModel) (*domain.User, error) {
	email, err := domain.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	role, err := domain.NewRole(m.Role)
	if err != nil {
		return nil, err
	}

	dni, err := domain.NewDni(m.Dni)
	if err != nil {
		return nil, err
	}

	return domain.Rehydrate(m.ID, m.FullName, email, m.Password, dni, role, m.Active, m.CreatedAt, m.UpdatedAt), nil
}
