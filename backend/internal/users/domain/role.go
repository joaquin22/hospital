package domain

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleStaff    Role = "staff"    // médico, encargado de sala, etc.
	RoleCustomer Role = "customer" // paciente, cliente que reserva
)
