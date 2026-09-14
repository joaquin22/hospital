package persistence

import (
	"github.com/joaquin22/hospital-api/internal/doctors/domain"
	"gorm.io/gorm"
)

type GormDoctorRepository struct {
	db *gorm.DB
}

func NewGormDoctorRepository(db *gorm.DB) *GormDoctorRepository {
	return &GormDoctorRepository{db: db}
}

func (r *GormDoctorRepository) Save(doctor *domain.Doctor) error {
	model := &DoctorModel{
		ID:            doctor.ID(),
		UserID:        doctor.UserID(),
		SpecialtyID:   doctor.SpecialityID(),
		LicenseNumber: doctor.LicenseNumber(),
		Active:        doctor.IsActive(),
		CreatedAt:     doctor.CreatedAt(),
		UpdatedAt:     doctor.UpdatedAt(),
	}

	if err := r.db.Save(model).Error; err != nil {
		return err
	}

	doctor.SyncPersisted(model.ID, model.CreatedAt, model.UpdatedAt)

	return nil
}

func (r *GormDoctorRepository) FindByID(id uint) (*domain.Doctor, error) {
	var model DoctorModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDoctorNotFound
		}
		return nil, err
	}

	return domain.Rehydrate(
		model.ID,
		model.UserID,
		model.SpecialtyID,
		model.LicenseNumber,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
	), nil
}

func (r *GormDoctorRepository) FindByUserID(userID uint) (*domain.Doctor, error) {
	var model DoctorModel
	if err := r.db.First(&model, "user_id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDoctorNotFound
		}
		return nil, err
	}
	return domain.Rehydrate(
		model.ID,
		model.UserID,
		model.SpecialtyID,
		model.LicenseNumber,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
	), nil
}

func (r *GormDoctorRepository) FindAll() ([]*domain.Doctor, error) {
	var models []DoctorModel
	if err := r.db.Order("id ASC").Find(&models).Error; err != nil {
		return nil, err
	}
	doctors := make([]*domain.Doctor, 0, len(models))
	for _, model := range models {
		doctors = append(doctors, domain.Rehydrate(
			model.ID,
			model.UserID,
			model.SpecialtyID,
			model.LicenseNumber,
			model.Active,
			model.CreatedAt,
			model.UpdatedAt,
		))
	}
	return doctors, nil
}
