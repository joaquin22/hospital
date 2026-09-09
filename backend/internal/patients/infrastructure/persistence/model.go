package persistence

import "time"

type PatientModel struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Dni       string
	Email     string
	Phone     string
	Active    bool `gorm:"default:true" json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PatientModel) TableName() string { return "patients" }
