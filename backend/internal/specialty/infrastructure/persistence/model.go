package persistence

import "time"

type SpecialityModel struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (SpecialityModel) TableName() string { return "specialties" }
