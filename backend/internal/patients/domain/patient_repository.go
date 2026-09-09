package domain

type PatientRepository interface {
	Save(patient *Patient) error
	FindByID(id uint) (*Patient, error)
	FindByDNI(dni Dni) (*Patient, error)
	FindByEmail(email Email) (*Patient, error)
	Delete(patient *Patient) error
	Update(patient *Patient) error
}
