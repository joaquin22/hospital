package domain

import "time"

type Speciality struct {
	id          uint
	name        string
	description string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewSpeciality(id uint, name, description string) (*Speciality, error) {
	return &Speciality{
		id:          id,
		name:        name,
		description: description,
		active:      true,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (s *Speciality) UpdateSpeciality(name, description string) error {
	s.name = name
	s.description = description
	s.updatedAt = time.Now()
	return nil
}

func Rehydrate(id uint, name, description string, active bool, createdAt, updatedAt time.Time) *Speciality {
	return &Speciality{
		id:          id,
		name:        name,
		description: description,
		active:      active,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (s *Speciality) SyncPersisted(id uint, createdAt, updatedAt time.Time) {
	if s.id == 0 {
		s.id = id
	}
	s.createdAt = createdAt
	s.updatedAt = updatedAt
}

func (s *Speciality) Activate() {
	s.active = true
}

func (s *Speciality) Deactivate() {
	s.active = false
}
func (s *Speciality) IsActive() bool {
	return s.active
}

func (s *Speciality) ID() uint             { return s.id }
func (s *Speciality) Name() string         { return s.name }
func (s *Speciality) Description() string  { return s.description }
func (s *Speciality) Active() bool         { return s.active }
func (s *Speciality) CreatedAt() time.Time { return s.createdAt }
func (s *Speciality) UpdatedAt() time.Time { return s.updatedAt }
