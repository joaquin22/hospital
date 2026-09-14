package domain

import "time"

type User struct {
	id        uint
	firstName string
	lastName  string
	email     Email
	password  string
	dni       Dni
	role      Role
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(firstName, lastName string, email Email, password string, dni Dni, role Role) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, ErrInvalidName
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
		dni:       dni,
		role:      role,
		active:    true,
	}, nil
}

func Rehydrate(id uint, firstName, lastName string, email Email, passwordHash string, dni Dni, role Role, active bool, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  passwordHash,
		dni:       dni,
		role:      role,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) SyncPersisted(id uint, createdAt, updatedAt time.Time) {
	if u.id == 0 {
		u.id = id
	}
	u.createdAt = createdAt
	u.updatedAt = updatedAt
}

func (u *User) Activate() {
	u.active = true
}

func (u *User) Deactivate() {
	u.active = false
}

func (u *User) ID() uint             { return u.id }
func (u *User) FirstName() string    { return u.firstName }
func (u *User) LastName() string     { return u.lastName }
func (u *User) Email() Email         { return u.email }
func (u *User) PasswordHash() string { return u.password }
func (u *User) Role() Role           { return u.role }
func (u *User) Dni() Dni             { return u.dni }
func (u *User) Active() bool         { return u.active }
func (u *User) CreatedAt() time.Time { return u.createdAt }
func (u *User) UpdatedAt() time.Time { return u.updatedAt }
