package domain

type DoctorRepository interface {
	Save(doctor *Doctor) error
	FindByID(id uint) (*Doctor, error)
	FindByUserID(userID uint) (*Doctor, error)
	FindAll() ([]*Doctor, error)
}
