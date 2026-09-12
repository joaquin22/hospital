package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	PatientPersistence "github.com/joaquin22/hospital-api/internal/patients/infrastructure/persistence"
	SpecialityPersistence "github.com/joaquin22/hospital-api/internal/specialty/infrastructure/persistence"
	UserPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
)

func ConnectDatabase(dsn string) (*gorm.DB, error) {

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&UserPersistence.UserModel{},
		&PatientPersistence.PatientModel{},
		&SpecialityPersistence.SpecialityModel{},
	)
}
