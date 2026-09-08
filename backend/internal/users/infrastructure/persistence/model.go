package persistence

import "time"

type UserModel struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"size:150;not null" json:"full_name"`
	Email     string `gorm:"size:100;unique;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Role      string
	Active    bool `gorm:"default:true" json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string { return "users" }
