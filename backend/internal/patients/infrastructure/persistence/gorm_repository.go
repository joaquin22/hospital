package persistence

import (
	"errors"

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

	patient.SyncPersisted(model.ID, model.CreatedAt, model.UpdatedAt)

	return nil
}

func (r *GormPatientRepository) FindByID(id string) (*domain.Patient, error) {
	var model PatientModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPatientNotFound
		}
		return nil, err
	}

	patient, err := toDomain(&model)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

func toDomain(model *PatientModel) (*domain.Patient, error) {

	email, err := domain.NewEmail(model.Email)
	if err != nil {
		return nil, err
	}

	dni, err := domain.NewDni(model.Dni)
	if err != nil {
		return nil, err
	}

	phone, err := domain.NewPhone(model.Phone)
	if err != nil {
		return nil, err
	}

	return domain.Rehydrate(
		model.ID,
		model.FirstName,
		model.LastName,
		dni,
		email,
		phone,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
	), nil
}
