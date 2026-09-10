package persistence

import (
	"errors"
	"fmt"

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

func (r *GormPatientRepository) FindByID(id uint) (*domain.Patient, error) {
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

func (r *GormPatientRepository) FindByEmail(email domain.Email) (*domain.Patient, error) {
	var model PatientModel

	if err := r.db.First(&model, "email = ?", email.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	patient, err := toDomain(&model)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *GormPatientRepository) FindByDni(dni domain.Dni) (*domain.Patient, error) {
	var model PatientModel
	if err := r.db.First(&model, "dni = ?", dni.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	patient, err := toDomain(&model)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *GormPatientRepository) ListPatients() ([]*domain.Patient, error) {
	var models []PatientModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	patients := make([]*domain.Patient, 0, len(models))
	for _, patient := range models {
		dni, err := domain.NewDni(patient.Dni)
		if err != nil {
			return nil, err
		}
		email, err := domain.NewEmail(patient.Email)
		if err != nil {
			return nil, err
		}

		phone, err := domain.NewPhone(patient.Phone)
		if err != nil {
			return nil, err
		}
		patients = append(patients, domain.Rehydrate(
			patient.ID, patient.FirstName, patient.LastName, dni, email, phone, patient.Active, patient.CreatedAt, patient.UpdatedAt,
		))
	}
	fmt.Println(patients)
	return patients, nil
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
