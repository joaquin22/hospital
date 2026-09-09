package domain

type UserRepository interface {
	// Define the methods that the UserRepository should have, for example:
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email Email) (*User, error)
	FindByDni(dni Dni) (*User, error)
}
