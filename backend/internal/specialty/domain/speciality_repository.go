package domain

type SpecialityRepository interface {
	Save(s *Speciality) error
	FindByID(id uint) (*Speciality, error)
	FindAll() ([]*Speciality, error)
}
