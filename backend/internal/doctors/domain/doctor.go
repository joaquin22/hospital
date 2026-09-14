package domain

import "time"

type Doctor struct {
	id            uint
	userId        uint
	specialityId  uint
	licenseNumber string
	active        bool
	createdAt     time.Time
	updatedAt     time.Time
}

func NewDoctor(userId, specialityId uint, licenseNumber string) (*Doctor, error) {
	return &Doctor{
		userId:        userId,
		specialityId:  specialityId,
		licenseNumber: licenseNumber,
		active:        true,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

func (d *Doctor) UpdateDoctor(specialityId uint, licenseNumber string) error {
	d.specialityId = specialityId
	d.licenseNumber = licenseNumber
	d.updatedAt = time.Now()
	return nil
}

func Rehydrate(id, userId, specialityId uint, licenseNumber string, active bool, createdAt, updatedAt time.Time) *Doctor {
	return &Doctor{
		id:            id,
		userId:        userId,
		specialityId:  specialityId,
		licenseNumber: licenseNumber,
		active:        active,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func (d *Doctor) SyncPersisted(id uint, createdAt, updatedAt time.Time) {
	if d.id == 0 {
		d.id = id
	}
	d.createdAt = createdAt
	d.updatedAt = updatedAt
}

func (d *Doctor) Activate() {
	d.active = true
}

func (d *Doctor) Deactivate() {
	d.active = false
}

func (d *Doctor) IsActive() bool {
	return d.active
}

func (d *Doctor) ID() uint              { return d.id }
func (d *Doctor) UserID() uint          { return d.userId }
func (d *Doctor) SpecialityID() uint    { return d.specialityId }
func (d *Doctor) LicenseNumber() string { return d.licenseNumber }
func (d *Doctor) CreatedAt() time.Time  { return d.createdAt }
func (d *Doctor) UpdatedAt() time.Time  { return d.updatedAt }
