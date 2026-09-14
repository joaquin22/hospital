package http

type CreateDoctorRequest struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Dni           string `json:"dni" binding:"required"`
	Role          string `json:"role" binding:"required"`
	SpecialtyID   uint   `json:"specialty_id" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
}
