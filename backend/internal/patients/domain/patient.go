package domain

import "time"

type Patient struct {
	id        uint
	firstName string
	lastName  string
	dni       Dni
	email     Email
	phone     Phone
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func NewPatient(firstName, lastName string, dni Dni, email Email, phone Phone) (*Patient, error) {
	return &Patient{
		firstName: firstName,
		lastName:  lastName,
		dni:       dni,
		email:     email,
		phone:     phone,
		active:    true,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func Rehydrate(id uint, firstName, lastName string, dni Dni, email Email, phone Phone, active bool, createdAt, updatedAt time.Time) *Patient {
	return &Patient{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		dni:       dni,
		email:     email,
		phone:     phone,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (p *Patient) SyncPersisted(id uint, createdAt, updatedAt time.Time) {
	if p.id == 0 {
		p.id = id
	}
	p.createdAt = createdAt
	p.updatedAt = updatedAt
}

func (p *Patient) Activate() {
	p.active = true
}

func (p *Patient) Deactivate() {
	p.active = false
}

func (p *Patient) IsActive() bool {
	return p.active
}

func (p *Patient) ID() uint             { return p.id }
func (p *Patient) FirstName() string    { return p.firstName }
func (p *Patient) LastName() string     { return p.lastName }
func (p *Patient) DNI() Dni             { return p.dni }
func (p *Patient) Email() Email         { return p.email }
func (p *Patient) Phone() Phone         { return p.phone }
func (p *Patient) CreatedAt() time.Time { return p.createdAt }
func (p *Patient) UpdatedAt() time.Time { return p.updatedAt }
