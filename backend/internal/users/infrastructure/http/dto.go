package http

type registerRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Dni      string `json:"dni" binding:"required"`
	Role     string `json:"role" binding:"required"` // ADMIN o RECEPTIONIST
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
