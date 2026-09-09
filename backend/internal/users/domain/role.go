package domain

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleStaff    Role = "staff"    // médico, encargado de sala, etc.
	RoleCustomer Role = "customer" // paciente, cliente que reserva
)

func NewRole(value string) (Role, error) {
	switch Role(value) {
	case RoleAdmin, RoleStaff, RoleCustomer:
		return Role(value), nil
	default:
		return "", ErrInvalidRole
	}
}
