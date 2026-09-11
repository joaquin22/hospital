package http

type createPatientRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UpdatePatientRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type PatchPatientRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Dni       *string `json:"dni"`
	Phone     *string `json:"phone"`
}
