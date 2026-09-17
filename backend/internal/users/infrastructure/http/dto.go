package http

type registerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Role      string `json:"role" binding:"required"` // ADMIN o RECEPTIONIST
}

type updateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	Role      string `json:"role" binding:"required"` // ADMIN o RECEPTIONIST
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type patchUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Dni       *string `json:"dni"`
	Role      *string `json:"role"` // ADMIN o RECEPTIONIST
}
