package persistence

import "time"

type UserModel struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`
	Email     string `gorm:"size:100;unique;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Dni       string `gorm:"size:20;unique;not null" json:"dni"`
	Role      string
	Active    bool `gorm:"default:true" json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string { return "users" }
