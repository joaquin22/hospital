package domain

type PatientRepository interface {
	Save(patient *Patient) error
	FindByID(id uint) (*Patient, error)
	FindByDni(dni Dni) (*Patient, error)
	FindByEmail(email Email) (*Patient, error)
	FindAll() ([]*Patient, error)
}
