package application

import (
	"time"

	"github.com/joaquin22/hospital-api/internal/specialty/domain"
)

type CreateSpecialityInput struct {
	Name        string
	Description string
}

type UpdateSpecialityInput struct {
	ID          uint
	Name        string
	Description string
}

type PatchSpecialityInput struct {
	ID          uint
	Name        *string
	Description *string
}

type SpecialityOutput struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toOutput(s *domain.Speciality) *SpecialityOutput {
	return &SpecialityOutput{
		ID:          s.ID(),
		Name:        s.Name(),
		Description: s.Description(),
		Active:      s.Active(),
		CreatedAt:   s.CreatedAt(),
		UpdatedAt:   s.UpdatedAt(),
	}
}
