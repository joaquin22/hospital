package persistence

import (
	"time"

	UserPersistence "github.com/joaquin22/hospital-api/internal/users/infrastructure/persistence"
)

type DoctorModel struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"uniqueIndex;not null"`

	SpecialtyID   uint   `gorm:"not null"`
	LicenseNumber string `gorm:"size:30;uniqueIndex;not null"`

	Active    bool `gorm:"default:true" json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User UserPersistence.UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (DoctorModel) TableName() string {
	return "doctors"
}
