package persistence

import (
	"errors"

	"github.com/joaquin22/hospital-api/internal/specialty/domain"
	"gorm.io/gorm"
)

type GormSpecialtyRepository struct {
	db *gorm.DB
}

func NewGormSpecialityRepository(db *gorm.DB) *GormSpecialtyRepository {
	return &GormSpecialtyRepository{db: db}
}

func (r *GormSpecialtyRepository) Save(speciality *domain.Speciality) error {
	model := &SpecialityModel{
		ID:          speciality.ID(),
		Name:        speciality.Name(),
		Description: speciality.Description(),
		Active:      speciality.Active(),
		CreatedAt:   speciality.CreatedAt(),
		UpdatedAt:   speciality.UpdatedAt(),
	}

	if err := r.db.Save(model).Error; err != nil {
		return err
	}

	speciality.SyncPersisted(model.ID, model.CreatedAt, model.UpdatedAt)
	return nil
}

func (r *GormSpecialtyRepository) FindByID(id uint) (*domain.Speciality, error) {
	var model SpecialityModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrSpecialityNotFound
		}
		return nil, err
	}
	return domain.Rehydrate(
		model.ID,
		model.Name,
		model.Description,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
	), nil
}

func (r *GormSpecialtyRepository) FindAll() ([]*domain.Speciality, error) {
	var models []SpecialityModel

	if err := r.db.Order("id ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	specialities := make([]*domain.Speciality, 0, len(models))

	for _, speciality := range models {
		specialities = append(specialities, domain.Rehydrate(
			speciality.ID,
			speciality.Name,
			speciality.Description,
			speciality.Active,
			speciality.CreatedAt,
			speciality.UpdatedAt,
		))
	}

	return specialities, nil
}
