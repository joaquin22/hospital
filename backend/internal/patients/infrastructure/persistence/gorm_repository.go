package persistence

import (
	"github.com/joaquin22/hospital-api/internal/patients/domain"
	"gorm.io/gorm"
)

type GormPatientRepository struct {
	db *gorm.DB
}

func NewGormPatientRepository(db *gorm.DB) *GormPatientRepository {
	return &GormPatientRepository{db: db}
}

func (r *GormPatientRepository) Save(patient *domain.Patient) error {
	model := &PatientModel{
		ID:        patient.ID(),
		FirstName: patient.FirstName(),
		LastName:  patient.LastName(),
		Email:     patient.Email().String(),
		Dni:       patient.Dni().String(),
		Phone:     patient.Phone().String(),
		Active:    patient.IsActive(),
		CreatedAt: patient.CreatedAt(),
		UpdatedAt: patient.UpdatedAt(),
	}
	if err := r.db.Save(model).Error; err != nil {
		return err
	}
	return nil
}
