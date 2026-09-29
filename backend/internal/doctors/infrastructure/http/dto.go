package http

type CreateDoctorRequest struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Dni           string `json:"dni" binding:"required"`
	Role          string `json:"role" binding:"required"`
	SpecialtyID   uint   `json:"speciality_id" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
}

type UpdateDoctorRequest struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Dni           string `json:"dni" binding:"required"`
	Role          string `json:"role" binding:"required"`
	SpecialtyID   uint   `json:"speciality_id" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
	Active        *bool  `json:"active"`
}

type PatchDoctorRequest struct {
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	Dni           *string `json:"dni"`
	Role          *string `json:"role"`
	SpecialtyID   *uint   `json:"speciality_id"`
	LicenseNumber *string `json:"license_number"`
	// Un doctor tiene dos banderas `active`: la de su ficha (doctors.active) y
	// la de su cuenta (users.active). Este único campo las actualiza a las dos
	// para que no puedan quedar contradictorias.
	Active *bool `json:"active"`
}
