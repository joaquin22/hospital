package persistence

import (
	"time"

	DoctorPersistence "github.com/joaquin22/hospital-api/internal/doctors/infrastructure/persistence"
)

type SpecialityModel struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Doctors []DoctorPersistence.DoctorModel `gorm:"foreignKey:SpecialtyID"`
}

func (SpecialityModel) TableName() string { return "specialties" }
