package domain

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email Email) (*User, error)
	FindByDni(dni Dni) (*User, error)
}
